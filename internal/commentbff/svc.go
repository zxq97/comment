package commentbff

import (
	"context"
	"time"

	"github.com/zxq97/comment/internal/commentsvc"
	"github.com/zxq97/comment/internal/env"
	"github.com/zxq97/relation/pkg/relation"
	"google.golang.org/grpc"
)

var (
	client commentsvc.CommentSvcClient
)

type CommentBff struct {
}

func InitCommentBff(conf *CommentBffConfig, conn, relationConn *grpc.ClientConn) error {
	err := env.InitLog(conf.LogPath)
	if err != nil {
		return err
	}
	client = commentsvc.NewCommentSvcClient(conn)
	relation.InitClient(relationConn)
	return nil
}

func (CommentBff) CreateComment(ctx context.Context, uid, auid int64) {
	res, err := relation.GetRelation(ctx, uid, []int64{auid}, relation.Source_APIGateway)
	if err != nil {

	}
	val, ok := res[auid]
	if !ok || time.Now().UnixMilli()-val.FollowTime < 100 {

	}
}
