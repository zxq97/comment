package commenttask

import (
	"github.com/zxq97/comment/internal/cache"
	"github.com/zxq97/comment/internal/constant"
	"github.com/zxq97/comment/internal/env"
	"github.com/zxq97/comment/internal/store"
	"github.com/zxq97/gotool/config"
	"github.com/zxq97/gotool/kafka"
)

var (
	consumers = []*kafka.Consumer{}
)

func InitCommentTask(conf *CommentTaskConfig) error {
	err := env.InitLog(conf.LogPath)
	if err != nil {
		return err
	}
	cache.InitCache(conf.Redis["redis"], conf.MC["mc"])
	err = store.InitStore(conf.Mysql["comment"])
	return err
}

func InitConsumer(conf *config.KafkaConf) error {
	publishConsumer, err := kafka.InitConsumer(conf.Addr, []string{constant.TopicCommentPublish}, "comment_task_publish", env.ApiLogger, env.ExcLogger)
	if err != nil {
		return err
	}
	publishConsumer.Start(publish)
	rebuildConsumer, err := kafka.InitConsumer(conf.Addr, []string{constant.TopicCommentCacheRebuild}, "comment_task_rebuild", env.ApiLogger, env.ExcLogger)
	if err != nil {
		return err
	}
	rebuildConsumer.Start(rebuild)
	attrConsumer, err := kafka.InitConsumer(conf.Addr, []string{constant.TopicCommentAttr}, "comment_task_attr", env.ApiLogger, env.ExcLogger)
	if err != nil {
		return err
	}
	attrConsumer.Start(attr)
	consumers = append(consumers, publishConsumer, rebuildConsumer, attrConsumer)
	return nil
}

func StopConsumer() {
	for _, v := range consumers {
		err := v.Stop()
		if err != nil {
			env.ExcLogger.Println("StopConsumer err", err)
		}
	}
}
