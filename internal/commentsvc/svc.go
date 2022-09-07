package commentsvc

import (
	"context"
	"fmt"

	"github.com/zxq97/comment/internal/cache"
	inconst "github.com/zxq97/comment/internal/constant"
	"github.com/zxq97/comment/internal/env"
	"github.com/zxq97/comment/internal/model"
	"github.com/zxq97/comment/internal/store"
	"github.com/zxq97/gotool/cast"
	"github.com/zxq97/gotool/kafka"
	"golang.org/x/sync/singleflight"
)

const (
	sfKeyGetTimeList    = "time_list_%d_%d_%d"        // objid objtype floor
	sfKeyGetTimeSubList = "time_sub_list_%d_%d_%d_%d" // objid objtype commentid floor
)

var (
	producer *kafka.Producer
	sfg      singleflight.Group
)

type CommentSvc struct {
}

func InitCommentSvc(conf *CommentSvcConfig) error {
	err := env.InitLog(conf.LogPath)
	if err != nil {
		return err
	}
	cache.InitCache(conf.Redis["redis"], conf.MC["mc"])
	err = store.InitStore(conf.Mysql["comment"])
	if err != nil {
		return err
	}
	producer, err = kafka.InitKafkaProducer(conf.Kafka["kafka"].Addr, env.ApiLogger, env.ExcLogger)
	return err
}

func (CommentSvc) CreateComment(ctx context.Context, req *CreateRequest) (*EmptyResponse, error) {
	bs, err := packKafkaMsg(ctx, req.Create, inconst.EventTypeCreate)
	if err != nil {
		return &EmptyResponse{}, err
	}
	shardKey := cast.FormatInt(req.Create.Uid)
	return &EmptyResponse{}, producer.SendMessage(inconst.TopicCommentPublish, []byte(shardKey), bs)
}

func (CommentSvc) ReplyComment(ctx context.Context, req *ReplyRequest) (*EmptyResponse, error) {
	bs, err := packKafkaMsg(ctx, req.Reply, inconst.EventTypeReply)
	if err != nil {
		return &EmptyResponse{}, err
	}
	shardKey := cast.FormatInt(req.Reply.Uid)
	return &EmptyResponse{}, producer.SendMessage(inconst.TopicCommentPublish, []byte(shardKey), bs)
}

func (CommentSvc) LikeComment(ctx context.Context, req *AttrRequest) (*EmptyResponse, error) {
	bs, err := packKafkaMsg(ctx, req, inconst.EventTypeLike)
	if err != nil {
		return &EmptyResponse{}, err
	}
	shardKey := cast.FormatInt(req.Uid)
	return &EmptyResponse{}, producer.SendMessage(inconst.TopicCommentAttr, []byte(shardKey), bs)
}

func (CommentSvc) HateComment(ctx context.Context, req *AttrRequest) (*EmptyResponse, error) {
	bs, err := packKafkaMsg(ctx, req, inconst.EventTypeHate)
	if err != nil {
		return &EmptyResponse{}, err
	}
	shardKey := cast.FormatInt(req.Uid)
	return &EmptyResponse{}, producer.SendMessage(inconst.TopicCommentAttr, []byte(shardKey), bs)
}

func (CommentSvc) GetCommentList(ctx context.Context, req *CommentListRequest) (*CommentListResponse, error) {
	var (
		roots    []int64
		subMap   map[int64][]int64
		floor    []*model.CommentFloor
		floorMap map[int64][]*model.CommentFloor
		err      error
	)
	sfk := fmt.Sprintf(sfKeyGetTimeList, req.List.ObjId, req.List.ObjType, req.List.Floor)
	_, err, _ = sfg.Do(sfk, func() (interface{}, error) {
		roots, err = cache.GetCommentListByTime(ctx, req.List.ObjId, req.List.Offset, req.List.Floor, int8(req.List.ObjType))
		if err != nil {
			env.ExcLogger.Printf("ctx %v GetCommentListByTime objid %v objtype %v floor %v err %v", req.List.ObjId, req.List.ObjType, req.List.Floor, err)
			bs, _ := packKafkaMsg(ctx, req.List, inconst.EventTypeListMissed)
			shareKey := cast.FormatInt(req.List.ObjId)
			_ = producer.SendMessage(inconst.TopicCommentCacheRebuild, []byte(shareKey), bs)
			floor, floorMap, err = store.GetCommentListByTime(ctx, req.List.ObjId, req.List.Offset, req.List.Floor, int8(req.List.ObjType))
			if err != nil {
				return nil, err
			}
			roots = floor2arr(floor...)
			subMap = floorMap2arr(floorMap)
		} else {
			subMap, err = cache.GetCommentsFirstSubListByTime(ctx, roots, req.List.ObjId, int8(req.List.ObjType))
			if err != nil {
				env.ExcLogger.Printf("ctx %v GetCommentSubListByTime objid %v objtype %v roots %v err %v", ctx, req.List.ObjId, req.List.ObjType, roots, err)
				bs, _ := packKafkaMsg(ctx, req.List, inconst.EventTypeListMissed)
				shareKey := cast.FormatInt(req.List.ObjId)
				_ = producer.SendMessage(inconst.TopicCommentCacheRebuild, []byte(shareKey), bs)
				floorMap, err = store.GetCommentsFirstSubListByTime(ctx, req.List.ObjId, int8(req.List.ObjType), roots)
				if err != nil {
					return nil, err
				}
				subMap = floorMap2arr(floorMap)
			}
		}
		return nil, nil
	})
	if err != nil {
		return &CommentListResponse{}, err
	}
	ids := make([]int64, 0, len(roots)<<2)
	for k, v := range subMap {
		ids = append(ids, k)
		ids = append(ids, v...)
	}
	itemMap := getCommentMetaCountData(ctx, req.List.ObjId, int8(req.List.ObjType), ids)
	itemList := make([]*CommentItem, 0, len(roots))
	listMap := make(map[int64]*ItemList, len(roots))
	for _, v := range roots {
		if item, ok := itemMap[v]; ok {
			itemList = append(itemList, item)
		}
		list := make([]*CommentItem, 0, inconst.SubCommentCount)
		for _, i := range subMap[v] {
			if item, ok := itemMap[i]; ok {
				list = append(list, item)
			}
		}
		listMap[v] = packItemList(list)
	}
	return &CommentListResponse{
		ItemList: packItemList(itemList),
		ItemMap:  listMap,
	}, nil
}

func (CommentSvc) GetCommentSubList(ctx context.Context, req *CommentListRequest) (*CommentSubListResponse, error) {
	var (
		ids   []int64
		floor []*model.CommentFloor
		err   error
	)
	sfk := fmt.Sprintf(sfKeyGetTimeSubList, req.List.ObjId, req.List.ObjType, req.List.Root, req.List.Floor)
	_, err, _ = sfg.Do(sfk, func() (interface{}, error) {
		ids, err = cache.GetCommentSubListByTime(ctx, req.List.ObjId, req.List.Root, req.List.Offset, req.List.Floor, int8(req.List.ObjType))
		if err != nil {
			env.ExcLogger.Printf("ctx %v GetCommentSubListByTime objid %v objtype %v commentid %v floor %v err %v", ctx, req.List.ObjId, req.List.ObjType, req.List.Root, req.List.Floor, err)
			bs, _ := packKafkaMsg(ctx, req.List, inconst.EventTypeSubListMissed)
			shareKey := cast.FormatInt(req.List.ObjId)
			_ = producer.SendMessage(inconst.TopicCommentCacheRebuild, []byte(shareKey), bs)
			floor, err = store.GetCommentSubListByTime(ctx, req.List.ObjId, req.List.Root, req.List.Offset, req.List.Floor, int8(req.List.ObjType))
			if err != nil {
				return nil, err
			}
			ids = floor2arr(floor...)
		}
		return nil, nil
	})
	if err != nil {
		return &CommentSubListResponse{}, nil
	}
	if req.List.Floor == 0 {
		ids = append(ids, req.List.Root)
	}
	itemMap := getCommentMetaCountData(ctx, req.List.ObjId, int8(req.List.ObjType), ids)
	itemList := make([]*CommentItem, 0, len(ids))
	root := &CommentItem{}
	for _, v := range ids {
		if v == req.List.Root {
			root = itemMap[req.List.Root]
			continue
		}
		if item, ok := itemMap[v]; ok {
			itemList = append(itemList, item)
		}
	}
	return &CommentSubListResponse{
		Item:     root,
		ItemList: packItemList(itemList),
	}, nil
}
