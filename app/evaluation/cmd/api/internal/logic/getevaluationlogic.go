package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEvaluationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEvaluationLogic {
	return &GetEvaluationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEvaluationLogic) GetEvaluation(req *types.GetEvaluationRequest) (resp *types.GetEvaluationResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
