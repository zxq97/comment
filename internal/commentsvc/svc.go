package commentsvc

import (
	"context"
	"fmt"

	"github.com/zxq97/comment/internal/cache"
	"github.com/zxq97/comment/internal/constant"
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

type ListResult struct {
	Roots  []int64
	SubMap map[int64][]int64
}

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
	bs, err := packKafkaMsg(ctx, req.Create, kafka.EventTypeCreate)
	if err != nil {
		return &EmptyResponse{}, err
	}
	shardKey := cast.FormatInt(req.Create.Uid)
	return &EmptyResponse{}, producer.SendMessage(kafka.TopicCommentPublish, []byte(shardKey), bs)
}

func (CommentSvc) ReplyComment(ctx context.Context, req *ReplyRequest) (*EmptyResponse, error) {
	bs, err := packKafkaMsg(ctx, req.Reply, kafka.EventTypeReply)
	if err != nil {
		return &EmptyResponse{}, err
	}
	shardKey := cast.FormatInt(req.Reply.Uid)
	return &EmptyResponse{}, producer.SendMessage(kafka.TopicCommentPublish, []byte(shardKey), bs)
}

func (CommentSvc) DeleteComment(ctx context.Context, req *DeleteRequest) (*EmptyResponse, error) {
	bs, err := packKafkaMsg(ctx, req, kafka.EventTypeDelete)
	if err != nil {
		return &EmptyResponse{}, err
	}
	shardKey := cast.FormatInt(req.CommentId)
	return &EmptyResponse{}, producer.SendMessage(kafka.TopicCommentOperator, []byte(shardKey), bs)
}

func (CommentSvc) GetCommentList(ctx context.Context, req *CommentListRequest) (*CommentListResponse, error) {
	var (
		floor    []*model.CommentFloor
		floorMap map[int64][]*model.CommentFloor
		val      interface{}
		err      error
	)
	sfk := fmt.Sprintf(sfKeyGetTimeList, req.List.ObjId, req.List.ObjType, req.List.Floor)
	val, err, _ = sfg.Do(sfk, func() (interface{}, error) {
		var (
			roots  []int64
			subMap map[int64][]int64
		)
		roots, err = cache.GetCommentListByTime(ctx, req.List.ObjId, req.List.Offset, req.List.Floor, int8(req.List.ObjType))
		if err != nil {
			env.ExcLogger.Printf("ctx %v GetCommentListByTime objid %v objtype %v floor %v err %v", req.List.ObjId, req.List.ObjType, req.List.Floor, err)
			bs, _ := packKafkaMsg(ctx, req.List, kafka.EventTypeListMissed)
			shareKey := cast.FormatInt(req.List.ObjId)
			_ = producer.SendMessage(kafka.TopicCommentCacheRebuild, []byte(shareKey), bs)
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
				bs, _ := packKafkaMsg(ctx, req.List, kafka.EventTypeListMissed)
				shareKey := cast.FormatInt(req.List.ObjId)
				_ = producer.SendMessage(kafka.TopicCommentCacheRebuild, []byte(shareKey), bs)
				floorMap, err = store.GetCommentsFirstSubListByTime(ctx, req.List.ObjId, int8(req.List.ObjType), roots)
				if err != nil {
					return nil, err
				}
				subMap = floorMap2arr(floorMap)
			}
		}
		return &ListResult{
			Roots:  roots,
			SubMap: subMap,
		}, nil
	})
	res, ok := val.(*ListResult)
	if err != nil || !ok {
		return &CommentListResponse{}, err
	}
	roots := res.Roots
	subMap := res.SubMap
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
		list := make([]*CommentItem, 0, constant.SubCommentCount)
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
		floor []*model.CommentFloor
		val   interface{}
		err   error
	)
	sfk := fmt.Sprintf(sfKeyGetTimeSubList, req.List.ObjId, req.List.ObjType, req.List.Root, req.List.Floor)
	val, err, _ = sfg.Do(sfk, func() (interface{}, error) {
		var ids []int64
		ids, err = cache.GetCommentSubListByTime(ctx, req.List.ObjId, req.List.Root, req.List.Offset, req.List.Floor, int8(req.List.ObjType))
		if err != nil {
			env.ExcLogger.Printf("ctx %v GetCommentSubListByTime objid %v objtype %v commentid %v floor %v err %v", ctx, req.List.ObjId, req.List.ObjType, req.List.Root, req.List.Floor, err)
			bs, _ := packKafkaMsg(ctx, req.List, kafka.EventTypeSubListMissed)
			shareKey := cast.FormatInt(req.List.ObjId)
			_ = producer.SendMessage(kafka.TopicCommentCacheRebuild, []byte(shareKey), bs)
			floor, err = store.GetCommentSubListByTime(ctx, req.List.ObjId, req.List.Root, req.List.Offset, req.List.Floor, int8(req.List.ObjType))
			if err != nil {
				return nil, err
			}
			ids = floor2arr(floor...)
		}
		return &ListResult{
			Roots: ids,
		}, nil
	})
	res, ok := val.(*ListResult)
	if err != nil || !ok {
		return &CommentSubListResponse{}, nil
	}
	ids := res.Roots
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
