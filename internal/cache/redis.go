package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/zxq97/comment/internal/constant"
	"github.com/zxq97/comment/internal/env"
	"github.com/zxq97/comment/internal/model"
	"github.com/zxq97/gotool/cast"
	"github.com/zxq97/gotool/concurrent"
)

const (
	commentCountField = "count"
	commentLikeField  = "like"
	commentHateField  = "field"

	commentCacheCountTTL        = 4 * time.Hour
	commentCacheListTTL         = 8 * time.Hour
	redisKeyZCommentTimeList    = "cmt_time_l_%d_%d"     // objid objtype
	redisKeyZCommentTimeSubList = "cmt_time_sl_%d_%d_%d" // objid objtype commentid
	redisKeyCommentSubjectCount = "cmt_s_cnt_%d_%d"      // objid objtype
	redisKeyHCommentCountData   = "cmt_cnt_%d_%d_%d"     // objid objtyoe commentid
)

func getCommentList(ctx context.Context, key string, offset, maxVal int32) ([]redis.Z, error) {
	var (
		zs  []redis.Z
		err error
	)
	if maxVal == 0 {
		zs, err = rdx.ZRevRangeWithScores(ctx, key, 0, int64(offset-1)).Result()
	} else {
		zs, err = rdx.ZRevRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{Max: "(" + strconv.Itoa(int(maxVal)), Count: int64(offset)}).Result()
	}
	if err != nil {
		return nil, err
	} else if len(zs) == 0 {
		return nil, redis.Nil
	}
	return zs, nil
}

func GetCommentListByTime(ctx context.Context, objid int64, offset, floor int32, objtype int8) ([]int64, error) {
	key := fmt.Sprintf(redisKeyZCommentTimeList, objid, objtype)
	zs, err := getCommentList(ctx, key, offset, floor)
	if err != nil {
		return nil, err
	}
	return zs2arr(zs), nil
}

func GetCommentSubListByTime(ctx context.Context, objid, commnentid int64, offset, floor int32, objtype int8) ([]int64, error) {
	key := fmt.Sprintf(redisKeyZCommentTimeSubList, objid, objtype, commnentid)
	zs, err := getCommentList(ctx, key, offset, floor)
	if err != nil {
		return nil, err
	}
	return zs2arr(zs), nil
}

func GetCommentsFirstSubListByTime(ctx context.Context, ids []int64, objid int64, objtype int8) (map[int64][]int64, error) {
	cmdMap := make(map[int64]*redis.ZSliceCmd, len(ids))
	pipe := rdx.Pipeline()
	for _, v := range ids {
		key := fmt.Sprintf(redisKeyZCommentTimeSubList, objid, objtype, v)
		cmdMap[v] = pipe.ZRevRangeWithScores(ctx, key, 0, constant.SubCommentCount-1)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}
	subMap := make(map[int64][]int64, len(ids))
	for k, v := range cmdMap {
		subMap[k] = zs2arr(v.Val())
	}
	return subMap, nil
}

func SetCommentTimeIndex(ctx context.Context, objid int64, objtype int8, floors []*model.CommentFloor, subMap map[int64][]*model.CommentFloor) error {
	wg := concurrent.NewWaitGroup()
	wg.Go(func() {
		key := fmt.Sprintf(redisKeyZCommentTimeList, objid, objtype)
		err := rdx.ZAddEX(ctx, key, floor2zs(floors...), commentCacheListTTL)
		if err != nil {
			env.ExcLogger.Printf("SetCommentTimeList objid %v objtype %v floor %#v err %v", objid, objtype, floors, err)
		}
	})
	for k, v := range subMap {
		cid, floor := k, v
		wg.Go(func() {
			key := fmt.Sprintf(redisKeyZCommentTimeSubList, objid, objtype, cid)
			err := rdx.ZAddEX(ctx, key, floor2zs(floor...), commentCacheListTTL)
			if err != nil {
				env.ExcLogger.Printf("SetCommentTimeSubList objid %v objtype %v cid %v floor %v err %v", objid, objtype, cid, floor, err)
			}
		})
	}
	wg.Wait()
	return nil
}

func addCommentTimeIndex(ctx context.Context, objid int64, objtype int8, floor *model.CommentFloor) error {
	key := fmt.Sprintf(redisKeyZCommentTimeList, objid, objtype)
	return rdx.ZAddXEX(ctx, key, floor2zs(floor), commentCacheListTTL)
}

func addCommentTimeSubIndex(ctx context.Context, objid int64, objtype int8, floor *model.CommentFloor) error {
	key := fmt.Sprintf(redisKeyZCommentTimeSubList, objid, objtype, floor.Root)
	return rdx.ZAddXEX(ctx, key, floor2zs(floor), commentCacheListTTL)
}

func addCommentCount(ctx context.Context, objid int64, objtype int8, floor *model.CommentFloor) {
	key := fmt.Sprintf(redisKeyCommentSubjectCount, objid, objtype)
	rkey := fmt.Sprintf(redisKeyHCommentCountData, objid, objtype, floor.Root)
	ckey := fmt.Sprintf(redisKeyHCommentCountData, objid, objtype, floor.ID)
	wg := concurrent.NewWaitGroup()
	wg.Go(func() {
		err := rdx.IncrByXEX(ctx, key, 1, commentCacheCountTTL)
		if err != nil {
			env.ExcLogger.Printf("ctx %v addcommentcount objid %v objtype %v err %v", ctx, objid, objtype, err)
		}
	})
	wg.Go(func() {
		err := rdx.HIncrByXEX(ctx, rkey, commentCountField, 1, commentCacheCountTTL)
		if err != nil {
			env.ExcLogger.Printf("ctx %v addcommentcount objid %v objtype %v root %v err %v", ctx, objid, objtype, floor.Root, err)
		}
	})
	wg.Go(func() {
		err := rdx.HIncrByXEX(ctx, ckey, commentCountField, 1, commentCacheCountTTL)
		if err != nil {
			env.ExcLogger.Printf("ctx %v addcommentcount objid %v objtype %v id %v err %v", ctx, objid, objtype, floor.ID, err)
		}
	})
	wg.Wait()
}

func CommentCreate(ctx context.Context, objid int64, objtype int8, floor *model.CommentFloor) error {
	err := addCommentTimeIndex(ctx, objid, objtype, floor)
	if err != nil {
		return err
	}
	addCommentCount(ctx, objid, objtype, floor)
	return nil
}

func CommentReply(ctx context.Context, objid int64, objtype int8, floor *model.CommentFloor) error {
	err := addCommentTimeSubIndex(ctx, objid, objtype, floor)
	if err != nil {
		return err
	}
	addCommentCount(ctx, objid, objtype, floor)
	return nil
}

func GetCommentsCountData(ctx context.Context, objid int64, objtype int8, ids []int64) (map[int64]*model.CommentCountData, []int64, error) {
	cmdMap := make(map[int64]*redis.StringStringMapCmd, len(ids))
	pipe := rdx.Pipeline()
	for _, v := range ids {
		key := fmt.Sprintf(redisKeyHCommentCountData, objid, objtype, v)
		cmdMap[v] = pipe.HGetAll(ctx, key)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return nil, ids, err
	}
	cntMap := make(map[int64]*model.CommentCountData, len(ids))
	missed := make([]int64, 0, len(ids))
	for k, v := range cmdMap {
		cd := &model.CommentCountData{}
		flag := false
		for f, s := range v.Val() {
			cnt := cast.Atoi(s, 0)
			flag = true
			switch f {
			case commentCountField:
				cd.Count = cnt
			case commentLikeField:
				cd.Like = cnt
			case commentHateField:
				cd.Hate = cnt
			}
		}
		if flag {
			cntMap[k] = cd
		}
	}
	for _, v := range ids {
		if _, ok := cntMap[v]; !ok {
			missed = append(missed, v)
		}
	}
	return cntMap, missed, nil
}

func SetCommentsCountData(ctx context.Context, objid int64, objtype int8, cntMap map[int64]*model.CommentCountData) {
	for k, v := range cntMap {
		key := fmt.Sprintf(redisKeyHCommentCountData, objid, objtype, k)
		fieldMap := map[string]interface{}{
			commentCountField: v.Count,
			commentLikeField:  v.Like,
			commentHateField:  v.Hate,
		}
		err := rdx.HMSetEX(ctx, key, fieldMap, commentCacheCountTTL)
		if err != nil {
			env.ExcLogger.Println("ctx %v SetCommentsCountData objid %v objtype %v id %v err %v", ctx, objid, objtype, k, err)
		}
	}
}
