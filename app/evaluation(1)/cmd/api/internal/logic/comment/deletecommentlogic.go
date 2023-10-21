package comment

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ctxdata"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/xerr"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation(1)/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation(1)/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *types.DeleteCommentReq) (resp *types.DeleteCommentRsp, err error) {

	comment, err := l.svcCtx.CommentModel.FindOne(l.ctx, int64(req.ObjectCommentId))

	if err != nil {
		return nil, xerr.NewErrCode(xerr.ErrGetOneComment)
	}

	if comment.Sid != ctxdata.GetStudentIdFromCtx(l.ctx) {
		return nil, xerr.NewErrCode(xerr.ErrPermissionDenied)
	}

	err = l.svcCtx.CommentModel.Delete(l.ctx, int64(req.ObjectCommentId))

	if err != nil {
		return nil, xerr.NewErrCode(xerr.ErrDeleteComment)
	}

	return &types.DeleteCommentRsp{
		Flag: true,
	}, nil
}
