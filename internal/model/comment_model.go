package model

import "time"

type SubjectRootCount struct {
	RootCount int32 `json:"root_count" db:"root_count"`
}

type IndexCount struct {
	Count int32 `json:"count" db:"count"`
}

type CommentFloor struct {
	ID    int64 `json:"id" db:"id"`
	Root  int64 `json:"root" db:"root"`
	Floor int32 `json:"floor" db:"floor"`
}

type CommentSubject struct {
	ObjID      int64     `json:"obj_id" db:"obj_id"`
	ObjType    int8      `json:"obj_type" db:"obj_type"`
	UID        int64     `json:"uid" db:"uid"`
	Count      int32     `json:"count" db:"count"`
	RootCount  int32     `json:"root_count" db:"root_count"`
	State      int8      `json:"state" db:"state"`
	CreateTime time.Time `json:"create_time" db:"create_tine"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}

type CommentIndex struct {
	ID         int64     `json:"id" db:"id"`
	ObjID      int64     `json:"obj_id" db:"obj_id"`
	ObjType    int8      `json:"obj_type" db:"obj_type"`
	UID        int64     `json:"uid" db:"uid"`
	Root       int64     `json:"root" db:"root"`
	Parent     int64     `json:"parent" db:"parent"`
	Floor      int32     `json:"floor" db:"floor"`
	Count      int32     `json:"count" db:"count"`
	Like       int32     `json:"like" db:"like"`
	Hate       int32     `json:"hate" db:"hate"`
	State      int8      `json:"state" db:"state"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}

type CommentContent struct {
	CommentID  int64     `json:"comment_id" db:"comment_id"`
	UID        int64     `json:"uid" db:"uid"`
	IP         int32     `json:"ip" db:"ip"`
	Message    string    `json:"message" db:"message"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}
