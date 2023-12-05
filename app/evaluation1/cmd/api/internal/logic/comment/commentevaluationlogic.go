package comment

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

type CommentEvaluationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentEvaluationLogic {
	return &CommentEvaluationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentEvaluationLogic) CommentEvaluation(req *types.CommentEvaluationReq) (resp *types.CommentEvaluationRsp, err error) {

	sid := ctxdata.GetStudentIdFromCtx(l.ctx)
	_, err = l.svcCtx.CommentModel.Insert(l.ctx, &model.Comment{
		EvaluationPid: int64(req.ObjectEvaluationId),
		Sid:           sid,
		Text:          req.Content,
		Date:          time.Now().Unix(),
	})

	if err != nil {
		return nil, xerr.NewErrCode(xerr.ErrCreateComment)
	}

	return &types.CommentEvaluationRsp{
		Flag: true,
	}, nil
}
