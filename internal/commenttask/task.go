package commenttask

import (
	"github.com/zxq97/comment/internal/cache"
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
	publishConsumer, err := kafka.InitConsumer(conf.Addr, []string{kafka.TopicCommentPublish}, "comment_task_publish", publish, env.ApiLogger, env.ExcLogger)
	if err != nil {
		return err
	}
	rebuildConsumer, err := kafka.InitConsumer(conf.Addr, []string{kafka.TopicCommentCacheRebuild}, "comment_task_rebuild", rebuild, env.ApiLogger, env.ExcLogger)
	if err != nil {
		return err
	}
	operatorConsumer, err := kafka.InitConsumer(conf.Addr, []string{kafka.TopicCommentOperator}, "comment_task_operator", operator, env.ApiLogger, env.ExcLogger)
	if err != nil {
		return err
	}
	opusOpeConsumer, err := kafka.InitConsumer(conf.Addr, []string{kafka.TopicOpusOperator}, "comment_task_opus_ope", opusOpe, env.ApiLogger, env.ExcLogger)
	if err != nil {
		return err
	}
	consumers = append(consumers, publishConsumer, rebuildConsumer, operatorConsumer, opusOpeConsumer)
	StartConsumer(consumers)
	return nil
}

func StartConsumer(consumers []*kafka.Consumer) {
	for _, v := range consumers {
		v.Start()
	}
}

func StopConsumer() {
	for _, v := range consumers {
		err := v.Stop()
		if err != nil {
			env.ExcLogger.Println("StopConsumer err", err)
		}
	}
}
