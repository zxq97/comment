package constant

const (
	SubCommentCount = 3

	TopicCommentPublish      = "comment_publish"
	TopicCommentCacheRebuild = "comment_cache_rebuild"
	TopicCommentAttr         = "comment_attr"

	EventTypeCreate        = 1
	EventTypeReply         = 2
	EventTypeListMissed    = 3
	EventTypeSubListMissed = 4
	EventTypeLike          = 5
	EventTypeHate          = 6

	DefaultListMissOffset = 100
)
