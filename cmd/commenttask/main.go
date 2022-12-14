package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zxq97/comment/internal/commenttask"
	"github.com/zxq97/gotool/config"
)

var (
	confPath = flag.String("conf", "", "configuration file")
	conf     commenttask.CommentTaskConfig
)

func main() {
	flag.Parse()
	err := config.LoadYaml(*confPath, &conf)
	if err != nil {
		panic(err)
	}
	conf.Initialize()

	err = commenttask.InitCommentTask(&conf)
	if err != nil {
		panic(err)
	}
	err = commenttask.InitConsumer(conf.Kafka["kafka"])
	if err != nil {
		panic(err)
	}

	http.Handle("/metrics", promhttp.Handler())

	errCh := make(chan error, 1)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		errCh <- http.ListenAndServe(conf.HTTPBind, nil)
	}()

	select {
	case err = <-errCh:
		commenttask.StopConsumer()
		log.Println("commenttask stop err", errCh)
	case sign := <-sigCh:
		commenttask.StopConsumer()
		log.Println("commenttask stop sign", sign)
	}
}
