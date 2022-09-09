package cache

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/zxq97/comment/internal/env"
	"github.com/zxq97/comment/internal/model"
)

const (
	commentCacheContentTTL = 24 * 3600
	mcKeyCommentMetadata   = "comment_metadata_%d_%d_%d" // objid objtype commentid
)

func GetCommentsMetaData(ctx context.Context, objid int64, objtyoe int8, ids []int64) (map[int64]*model.CommentMetaData, []int64, error) {
	keys := make([]string, len(ids))
	for k, v := range ids {
		keys[k] = fmt.Sprintf(mcKeyCommentMetadata, objid, objtyoe, v)
	}
	val, err := mcx.GetMultiCtx(ctx, keys)
	if err != nil {
		return nil, nil, err
	}
	metaMap := make(map[int64]*model.CommentMetaData, len(ids))
	missed := make([]int64, 0, len(ids))
	for _, v := range val {
		meta := &model.CommentMetaData{}
		err = proto.Unmarshal(v.Value, meta)
		if err != nil {
			env.ExcLogger.Println("set metadata unmarshal err", v, err)
			continue
		}
		metaMap[meta.CommentId] = meta
	}
	for _, k := range ids {
		if _, ok := metaMap[k]; !ok {
			ids = append(missed, k)
		}
	}
	return metaMap, missed, nil
}

func SetCommentsMetaData(ctx context.Context, objid int64, objtype int8, metaMap map[int64]*model.CommentMetaData) {
	for k, v := range metaMap {
		key := fmt.Sprintf(mcKeyCommentMetadata, objid, objtype, k)
		bs, err := proto.Marshal(v)
		if err != nil {
			env.ExcLogger.Println("set metadata marshal err", v, err)
			continue
		}
		err = mcx.SetCtx(ctx, key, bs, commentCacheContentTTL)
		if err != nil {
			env.ExcLogger.Println("mcx set err", key, string(bs), err)
		}
	}
}

func deleteCommentMetaData(ctx context.Context, objid int64, objtype int8, ids []int64) {
	for _, v := range ids {
		key := fmt.Sprintf(mcKeyCommentMetadata, objid, objtype, v)
		err := mcx.Delete(key)
		if err != nil {
			env.ExcLogger.Println("mcx del err", key, err)
		}
	}
}
