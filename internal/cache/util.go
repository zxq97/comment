package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/zxq97/comment/internal/model"
)

func floor2zs(floors ...*model.CommentFloor) []*redis.Z {
	zs := make([]*redis.Z, len(floors))
	for k, v := range floors {
		zs[k] = &redis.Z{
			Member: v.ID,
			Score:  float64(v.Floor),
		}
	}
	return zs
}

func zs2arr(zs []redis.Z) []int64 {
	ids := make([]int64, 0, len(zs))
	for _, v := range zs {
		id, ok := v.Member.(int64)
		if ok {
			ids = append(ids, id)
		}
	}
	return ids
}
