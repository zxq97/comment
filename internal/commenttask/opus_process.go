package commenttask

import (
	"context"

	"github.com/zxq97/gotool/kafka"
)

func opusOpe(ctx context.Context, kfkmsg *kafka.KafkaMessage) {
	switch kfkmsg.EventType {
	case kafka.EventTypeDelete:

	}
}
