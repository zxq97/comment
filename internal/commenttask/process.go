package commenttask

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zxq97/comment/internal/cache"
	"github.com/zxq97/comment/internal/commentsvc"
	"github.com/zxq97/comment/internal/constant"
	"github.com/zxq97/comment/internal/env"
	"github.com/zxq97/comment/internal/model"
	"github.com/zxq97/comment/internal/store"
	"github.com/zxq97/gotool/kafka"
)

func publish(ctx context.Context, kfkmsg *kafka.KafkaMessage) {
	comment := &commentsvc.PublishRequest{}
	floor := &model.CommentFloor{}
	err := proto.Unmarshal(kfkmsg.Message, comment)
	if err != nil {
		env.ExcLogger.Printf("ctx %v publish Unmarshal kfkmsg %#v err %v", ctx, kfkmsg, err)
		return
	}
	switch kfkmsg.EventType {
	case constant.EventTypeCreate:
		floor, err = store.CommentCreate(ctx, comment.CommentId, comment.ObjId, comment.AuthorUid, comment.Uid, comment.Ip, int8(comment.ObjType), comment.Message)
		if err != nil {
			env.ExcLogger.Printf("ctx %v publish CommentCreate comment %#v err %v", ctx, comment, err)
			return
		}
		err = cache.CommentCreate(ctx, comment.ObjId, int8(comment.ObjType), floor)
		if err != nil {
			env.ExcLogger.Printf("ctx %v publish AddCommentTimeIndex comment %#v floor %#v err %v", ctx, comment, floor, err)
			return
		}
	case constant.EventTypeReply:
		floor, err = store.CommentReply(ctx, comment.CommentId, comment.ObjId, comment.Uid, comment.Root, comment.Parent, comment.Ip, int8(comment.ObjType), comment.Message)
		if err != nil {
			env.ExcLogger.Printf("ctx %v publish CommentReply comment %#v err %v", ctx, comment, err)
			return
		}
		err = cache.CommentReply(ctx, comment.ObjId, int8(comment.ObjType), floor)
		if err != nil {
			env.ExcLogger.Printf("ctx %v publish AddCommentTimeSubIndex comment %#v floor %#v err %v", ctx, comment, floor, err)
			return
		}
	}
}

func rebuild(ctx context.Context, kfkmsg *kafka.KafkaMessage) {
	var (
		list   *commentsvc.ListRequest
		floor  []*model.CommentFloor
		subMap map[int64][]*model.CommentFloor
	)
	err := proto.Unmarshal(kfkmsg.Message, list)
	if err != nil {
		env.ExcLogger.Printf("ctx %v publish Unmarshal kfkmsg %#v err %v", ctx, kfkmsg, err)
		return
	}
	ok := cache.CheckTimeIndexExist(ctx, list.ObjId, list.Root, kfkmsg.EventType, list.Floor, int8(list.ObjType))
	if ok {
		return
	}
	switch kfkmsg.EventType {
	case constant.EventTypeListMissed:
		floor, subMap, err = store.GetCommentListByTime(ctx, list.ObjId, constant.DefaultListMissOffset, list.Floor, int8(list.ObjType))
		if err != nil {
			env.ExcLogger.Printf("ctx %v get list objid %v objtype %v floor %v err %v", ctx, list.ObjId, list.ObjType, list.Floor, err)
			return
		}
		err = cache.SetCommentTimeIndex(ctx, list.ObjId, int8(list.ObjType), floor, subMap)
		if err != nil {
			env.ExcLogger.Printf("ctx %v set index list objid %v objtype %v floor %v err %v", ctx, list.ObjId, list.ObjType, floor)
			return
		}
	case constant.EventTypeSubListMissed:
		floor, err = store.GetCommentSubListByTime(ctx, list.ObjId, list.Root, list.Offset, list.Floor, int8(list.ObjType))
		if err != nil {
			env.ExcLogger.Printf("ctx %v get sub list objid %v objtype %v root %v floor %v err %v", list.ObjId, list.ObjType, list.Root, list.Floor, err)
			return
		}
		err = cache.SetCommentTimeIndex(ctx, list.ObjId, int8(list.ObjType), floor, nil)
		if err != nil {
			env.ExcLogger.Printf("ctx %v set sub list objid %v objtype %v root %v floor %v err %v", list.ObjId, list.ObjType, list.Root, list.Floor, err)
			return
		}
	}
}

func attr(ctx context.Context, kfkmsg *kafka.KafkaMessage) {

}
