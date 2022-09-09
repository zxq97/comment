package model

import "time"

type SubjectRootCount struct {
	RootCount int32 `db:"root_count"`
}

type IndexCount struct {
	Count int32 `db:"count"`
}

type CommentFloor struct {
	ID    int64 `db:"id"`
	Root  int64 `db:"root"`
	Floor int32 `db:"floor"`
}

type CommentSubject struct {
	ObjID      int64     `db:"obj_id"`
	ObjType    int8      `db:"obj_type"`
	UID        int64     `db:"uid"`
	Count      int32     `db:"count"`
	RootCount  int32     `db:"root_count"`
	State      int8      `db:"state"`
	CreateTime time.Time `db:"create_tine"`
	UpdateTime time.Time `db:"update_time"`
}

type CommentIndex struct {
	ID         int64     `db:"id"`
	ObjID      int64     `db:"obj_id"`
	ObjType    int8      `db:"obj_type"`
	UID        int64     `db:"uid"`
	Root       int64     `db:"root"`
	Parent     int64     `db:"parent"`
	Floor      int32     `db:"floor"`
	Count      int32     `db:"count"`
	Like       int32     `db:"like"`
	Hate       int32     `db:"hate"`
	State      int8      `db:"state"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}

type CommentContent struct {
	CommentID  int64     `db:"comment_id"`
	UID        int64     `db:"uid"`
	IP         int32     `db:"ip"`
	Message    string    `db:"message"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}
