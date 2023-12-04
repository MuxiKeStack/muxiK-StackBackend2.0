package reply

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation1/model"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ctxdata"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/xerr"
	"time"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation1/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation1/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReplyCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReplyCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReplyCommentLogic {
	return &ReplyCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReplyCommentLogic) ReplyComment(req *types.ReplyCommentReq) (resp *types.ReplyCommentRsp, err error) {

	sid := ctxdata.GetStudentIdFromCtx(l.ctx)
	_, err = l.svcCtx.ReplyModel.Insert(l.ctx, &model.Reply{
		CommentId: int64(req.ObjectCommentId),
		Sid:       sid,
		Text:      req.Content,
		Date:      time.Now().Unix(),
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ErrCreateReply)
	}

	return &types.ReplyCommentRsp{
		Flag: true,
	}, nil
}
