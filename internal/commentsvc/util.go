package commentsvc

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zxq97/comment/internal/cache"
	"github.com/zxq97/comment/internal/env"
	"github.com/zxq97/comment/internal/model"
	"github.com/zxq97/comment/internal/store"
	"github.com/zxq97/gotool/concurrent"
	"github.com/zxq97/gotool/constant"
	"github.com/zxq97/gotool/generate"
	"github.com/zxq97/gotool/kafka"
)

func packKafkaMsg(ctx context.Context, req proto.Message, eventtype int32) ([]byte, error) {
	trace, ok := ctx.Value(constant.TraceIDKey).(string)
	if !ok {
		trace = generate.UUIDStr()
	}
	bs, err := proto.Marshal(req)
	if err != nil {
		env.ExcLogger.Printf("ctx %v CreateComment Marshal req %#v err %v", ctx, req, err)
		return nil, err
	}
	kfkmsg := &kafka.KafkaMessage{
		TraceId:   trace,
		EventType: eventtype,
		Message:   bs,
	}
	bs, err = proto.Marshal(kfkmsg)
	if err != nil {
		env.ExcLogger.Printf("ctx %v CreateComment Marshal kfkmsg %#v err %v", ctx, kfkmsg, err)
	}
	return bs, err
}

func packCommentItem(metaMap map[int64]*model.CommentMetaData, cntMap map[int64]*model.CommentCountData) map[int64]*CommentItem {
	itemMap := make(map[int64]*CommentItem, len(metaMap))
	for k, v := range metaMap {
		itemMap[k].MetaData = v
		if data, ok := cntMap[k]; ok {
			itemMap[k].CntData = data
		}
	}
	return itemMap
}

func packItemList(list []*CommentItem) *ItemList {
	return &ItemList{
		ItemList: list,
	}
}

func floor2arr(floors ...*model.CommentFloor) []int64 {
	ids := make([]int64, 0, len(floors))
	for _, v := range floors {
		ids = append(ids, v.ID)
	}
	return ids
}

func floorMap2arr(m map[int64][]*model.CommentFloor) map[int64][]int64 {
	fm := make(map[int64][]int64, len(m))
	for k, v := range m {
		fm[k] = make([]int64, 0, len(v))
		for _, c := range v {
			fm[k] = append(fm[k], c.ID)
		}
	}
	return fm
}

func getCommentMetaCountData(ctx context.Context, objid int64, objtype int8, ids []int64) map[int64]*CommentItem {
	var (
		metaMap    map[int64]*model.CommentMetaData
		cntMap     map[int64]*model.CommentCountData
		metaDBMap  map[int64]*model.CommentMetaData
		cntDBMap   map[int64]*model.CommentCountData
		metaMissed []int64
		cntMissed  []int64
		err        error
	)
	wg := concurrent.NewWaitGroup()
	wg.Go(func() {
		metaMap, metaMissed, err = cache.GetCommentsMetaData(ctx, objid, objtype, ids)
		if err != nil || len(metaMissed) != 0 {
			metaDBMap, _, err = store.GetCommentItem(ctx, objid, objtype, metaMissed)
			if err != nil {
				env.ExcLogger.Printf("ctx %v GetCommentItem objid %v objtype %v ids %v err %v", ctx, objid, objtype, ids, err)
				return
			}
			cache.SetCommentsMetaData(ctx, objid, objtype, metaDBMap)
			for k, v := range metaDBMap {
				metaMap[k] = v
			}
		}
	})
	wg.Go(func() {
		cntMap, cntMissed, err = cache.GetCommentsCountData(ctx, objid, objtype, ids)
		if err != nil || len(cntMissed) != 0 {
			_, cntDBMap, err = store.GetCommentItem(ctx, objid, objtype, cntMissed)
			if err != nil {
				env.ExcLogger.Printf("ctx %v GetCommentItem objid %v objtype %v ids %v err %v", ctx, objid, objtype, cntMissed, err)
				return
			}
			cache.SetCommentsCountData(ctx, objid, objtype, cntDBMap)
			for k, v := range cntMap {
				cntMap[k] = v
			}
		}
	})
	wg.Wait()
	return packCommentItem(metaMap, cntMap)
}
