package store

import "github.com/zxq97/comment/internal/model"

func packComment(m map[int64]*model.CommentContent, idxMap map[int64]*model.CommentIndex) (map[int64]*model.CommentMetaData, map[int64]*model.CommentCountData) {
	metaMap := make(map[int64]*model.CommentMetaData, len(idxMap))
	cntMap := make(map[int64]*model.CommentCountData, len(idxMap))
	for k, v := range idxMap {
		if c, ok := m[k]; ok {
			metaMap[k] = &model.CommentMetaData{
				ObjId:      v.ObjID,
				ObjType:    int32(v.ObjType),
				Uid:        v.UID,
				CommentId:  k,
				Root:       v.Root,
				Parent:     v.Parent,
				Floor:      v.Floor,
				Ip:         c.IP,
				CreateTime: c.CreateTime.UnixMilli(),
				Message:    c.Message,
			}
			cntMap[k] = &model.CommentCountData{
				Like:  v.Like,
				Hate:  v.Hate,
				Count: v.Count,
			}
		}
	}
	return metaMap, cntMap
}
