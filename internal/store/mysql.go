package store

import (
	"context"
	"fmt"

	"github.com/zxq97/comment/internal/constant"
	"github.com/zxq97/comment/internal/model"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

const (
	tableCommentSubject = "comment_subject"
	tableCommentIndex   = "comment_index"
	tableCommentContent = "comment_content"
)

func addCommentContent(ctx context.Context, commentid, uid int64, ip int32, message string) error {
	_, err := dbCli.InsertInto(tableCommentContent).Columns("comment_id", "uid", "ip", "message").Values(commentid, uid, ip, message).ExecContext(ctx)
	return err
}

func CommentCreate(ctx context.Context, commentid, objid, authoruid, uid int64, ip int32, objtype int8, message string) (*model.CommentFloor, error) {
	err := addCommentContent(ctx, commentid, uid, ip, message)
	if err != nil {
		return nil, err
	}
	floor := &model.CommentFloor{}
	err = dbCli.Tx(ctx, func(sess sqlbuilder.Tx) error {
		sql := "INSERT INTO %s (`obj_id`, `obj_type`, `uid`, `count`, `root_count`) VALUES (?, ?, ?, 1, 1) ON DUPLICATE KEY UPDATE `count` = `count` + 1, `root_count` = `root_count` + 1"
		_, err = sess.Exec(fmt.Sprintf(sql, tableCommentSubject), objid, objtype, authoruid)
		if err != nil {
			return err
		}
		rootCnt := &model.SubjectRootCount{}
		filter := db.Cond{"obj_id": objid, "obj_type": objtype, "state": 0}
		err = sess.Select("root_count").From(tableCommentSubject).Where(filter).One(rootCnt)
		if err != nil {
			return err
		}
		floor.ID = commentid
		floor.Floor = rootCnt.RootCount
		_, err = sess.InsertInto(tableCommentIndex).Columns("id", "obj_id", "obj_type", "uid", "floor").Values(commentid, objid, objtype, uid, rootCnt.RootCount).Exec()
		return err
	})
	return floor, err
}

func CommentReply(ctx context.Context, commentid, objid, uid, root, parent int64, ip int32, objtype int8, message string) (*model.CommentFloor, error) {
	err := addCommentContent(ctx, commentid, uid, ip, message)
	if err != nil {
		return nil, err
	}
	floor := &model.CommentFloor{}
	err = dbCli.Tx(ctx, func(sess sqlbuilder.Tx) error {
		filter := db.Cond{"obj_id": objid, "obj_type": objtype, "state": 0}
		_, err = sess.Update(tableCommentSubject).Set("count = count + 1").Where(filter).Exec()
		if err != nil {
			return err
		}
		filter["id"] = root
		_, err = sess.Update(tableCommentIndex).Set("count = count + 1").Where(filter).Limit(1).Exec()
		if err != nil {
			return err
		}
		idx := &model.IndexCount{}
		err = sess.Select("count").From(tableCommentIndex).Where(filter).One(idx)
		if err != nil {
			return err
		}
		floor.ID = commentid
		floor.Root = root
		floor.Floor = idx.Count
		_, err = sess.InsertInto(tableCommentIndex).Columns("id", "obj_id", "obj_type", "uid", "root", "parent", "floor").Values(commentid, objid, objtype, uid, root, parent, idx.Count).Exec()
		return err
	})
	return floor, err
}

func getCommentContent(ctx context.Context, ids []int64) (map[int64]*model.CommentContent, error) {
	contents := []*model.CommentContent{}
	filter := db.Cond{"comment_id IN": ids}
	err := dbCli.WithContext(ctx).SelectFrom(tableCommentContent).Where(filter).All(&contents)
	if err != nil {
		return nil, err
	}
	m := make(map[int64]*model.CommentContent, len(ids))
	for _, v := range contents {
		m[v.CommentID] = v
	}
	return m, nil
}

func getCommentIndex(ctx context.Context, objid int64, objtype int8, ids []int64) (map[int64]*model.CommentIndex, error) {
	idxs := []*model.CommentIndex{}
	filter := db.Cond{"obj_id": objid, "obj_type": objtype, "id IN": idxs, "state": 0}
	err := dbCli.WithContext(ctx).SelectFrom(tableCommentIndex).Where(filter).All(&idxs)
	if err != nil {
		return nil, err
	}
	idxMap := make(map[int64]*model.CommentIndex, len(ids))
	for _, v := range idxs {
		idxMap[v.ID] = v
	}
	return idxMap, nil
}

func GetCommentListByTime(ctx context.Context, objid int64, offset, floor int32, objtype int8) ([]*model.CommentFloor, map[int64][]*model.CommentFloor, error) {
	floors := []*model.CommentFloor{}
	filter := db.Cond{"obj_id": objid, "obj_type": objtype, "state": 0, "root": 0, "parent": 0}
	if floor > 0 {
		filter["floor <"] = floor
	}
	err := dbCli.WithContext(ctx).Select("id", "floor").From(tableCommentIndex).Where(filter).OrderBy("floor DESC").Limit(int(offset)).All(&floors)
	if err != nil {
		return nil, nil, err
	}
	ids := make([]int64, 0, len(floors))
	for _, v := range floors {
		ids = append(ids, v.ID)
	}
	subFloors := []*model.CommentFloor{}
	filter = db.Cond{"obj_id": objid, "obj_type": objtype, "state": 0, "root IN": ids}
	err = dbCli.WithContext(ctx).Select("id", "root", "floor").From(tableCommentIndex).Where(filter).OrderBy("floor DESC").Limit(constant.SubCommentCount).All(&subFloors)
	if err != nil {
		return nil, nil, err
	}
	subMap, err := GetCommentsFirstSubListByTime(ctx, objid, objtype, ids)
	return floors, subMap, err
}

func GetCommentSubListByTime(ctx context.Context, objid, commentid int64, offset, floor int32, objtype int8) ([]*model.CommentFloor, error) {
	floors := []*model.CommentFloor{}
	filter := db.Cond{"obj_id": objid, "obj_type": objtype, "state": 0, "root": commentid}
	if floor > 0 {
		filter["floor <"] = floor
	}
	err := dbCli.WithContext(ctx).Select("id", "root", "floor").From(tableCommentIndex).Where(filter).OrderBy("floor DESC").Limit(int(offset)).All(&floors)
	return floors, err
}

func GetCommentsFirstSubListByTime(ctx context.Context, objid int64, objtype int8, ids []int64) (map[int64][]*model.CommentFloor, error) {
	floors := []*model.CommentFloor{}
	filter := db.Cond{"obj_id": objid, "obj_type": objtype, "state": 0, "root IN": ids}
	err := dbCli.WithContext(ctx).Select("id", "root", "floor").From(tableCommentIndex).Where(filter).OrderBy("floor DESC").Limit(constant.SubCommentCount).All(&floors)
	if err != nil {
		return nil, err
	}
	subMap := make(map[int64][]*model.CommentFloor, len(ids))
	for _, v := range floors {
		if subMap[v.Root] == nil {
			subMap[v.Root] = make([]*model.CommentFloor, 0, constant.SubCommentCount)
		}
		subMap[v.Root] = append(subMap[v.Root], v)
	}
	for _, k := range ids {
		if _, ok := subMap[k]; !ok {
			subMap[k] = make([]*model.CommentFloor, 1)
			subMap[k][0] = &model.CommentFloor{}
		}
	}
	return subMap, nil
}

func GetCommentItem(ctx context.Context, objid int64, objtype int8, ids []int64) (map[int64]*model.CommentMetaData, map[int64]*model.CommentCountData, error) {
	m, err := getCommentContent(ctx, ids)
	if err != nil {
		return nil, nil, err
	}
	idxMap, err := getCommentIndex(ctx, objid, objtype, ids)
	if err != nil {
		return nil, nil, err
	}
	metaMap, cntMap := packComment(m, idxMap)
	return metaMap, cntMap, nil
}
