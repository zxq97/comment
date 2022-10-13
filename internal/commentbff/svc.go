package commentbff

import (
	"github.com/zxq97/comment/internal/commentsvc"
	"github.com/zxq97/comment/internal/env"
	"google.golang.org/grpc"
)

var (
	client commentsvc.CommentSvcClient
)

type CommentBff struct {
}

func InitCommentBff(conf *CommentBffConfig, commentConn *grpc.ClientConn) error {
	err := env.InitLog(conf.LogPath)
	if err != nil {
		return err
	}
	client = commentsvc.NewCommentSvcClient(commentConn)
	return nil
}
