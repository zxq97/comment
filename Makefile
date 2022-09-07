GOCMD=go
GOBUILD=${GOCMD} build -mod=mod
GOCLEAN=${GOCMD} clean

build: commentsvc

.PHONY: \
    commentbff commentsvc commenttask commentadmin commentbffprod commentsvcprod commenttaskprod commentadminprod

clean:
	${GOCLEAN}

commentsvcprod:
	${GOBUILD} -o /home/work/run/comment_svc github.com/zxq97/comment/cmd/commentsvc

commentsvc:
	${GOBUILD} -o /Users/zongxingquan/goland/run/comment_svc github.com/zxq97/comment/cmd/commentsvc

commenttaskprod:
	${GOBUILD} -o /home/work/run/comment_task github.com/zxq97/comment/cmd/commenttask

commenttask:
	${GOBUILD} -o /Users/zongxingquan/goland/run/comment_task github.com/zxq97/comment/cmd/commenttask

commentbffprod:
	${GOBUILD} -o /home/work/run/comment_bff github.com/zxq97/comment/cmd/commentbff

commentbff:
	${GOBUILD} -o /Users/zongxingquan/goland/run/comment_bff github.com/zxq97/comment/cmd/commentbff

commentadminprod:
	${GOBUILD} -o /home/work/run/comment_admin github.com/zxq97/comment/cmd/commentadmin

commentadmin:
	${GOBUILD} -o /Users/zongxingquan/goland/run/comment_admin github.com/zxq97/comment/cmd/commentadmin
