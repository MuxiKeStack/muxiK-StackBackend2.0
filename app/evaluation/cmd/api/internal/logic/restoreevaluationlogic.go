package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"

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
	// todo: test
	response, err := l.svcCtx.InfoRpc.UpdateEvaluation(l.ctx, &pb.UpdateEvaluationRequest{E: &pb.Evaluation{
		Pid:     req.PostId,
		Deleted: false,
	}})
	if err != nil {
		resp.Code = 500
		resp.Msg = response.String()
	}
	return
}
