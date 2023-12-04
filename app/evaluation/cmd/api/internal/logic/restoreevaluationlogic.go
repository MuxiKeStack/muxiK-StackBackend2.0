package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RestoreEvaluationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRestoreEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RestoreEvaluationLogic {
	return &RestoreEvaluationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RestoreEvaluationLogic) RestoreEvaluation(req *types.RestoreEvaluationRequest) (resp *types.RestoreEvaluationResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
