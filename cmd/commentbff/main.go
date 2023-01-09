package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/zxq97/comment/internal/commentbff"
	"github.com/zxq97/gotool/config"
	"github.com/zxq97/gotool/rpc"
)

var (
	confPath = flag.String("conf", "", "configuration file")
	conf     commentbff.CommentBffConfig
)

func main() {
	flag.Parse()
	err := config.LoadYaml(*confPath, &conf)
	if err != nil {
		panic(err)
	}
	conf.Initialize()
	etcdCli, err := conf.Etcd["etcd"].InitEtcd()
	if err != nil {
		panic(err)
	}
	conn, err := rpc.NewGrpcConn(etcdCli, "commentsvc", conf.Hystrix["commentsvc"])
	if err != nil {
		panic(err)
	}
	relationConn, err := rpc.NewGrpcConn(etcdCli, "relationbff", conf.Hystrix["relationbff"])
	if err != nil {
		panic(err)
	}

	errCh := make(chan error, 1)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {

	}()
	go func() {
		errCh <- http.ListenAndServe(conf.Svc.HttpBind, nil)
	}()

	select {
	case err = <-errCh:
		log.Println("commentbff stop err", err)
	case sig := <-sigCh:
		log.Println("commenetbff stop sign", sig)
	}
}
