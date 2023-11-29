package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/types"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEvaluationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEvaluationLogic {
	return &DeleteEvaluationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteEvaluationLogic) DeleteEvaluation(req *types.DeleteEvaluationRequest) (resp *types.DeleteEvaluationResponse, err error) {
	// todo: test
	response, err := l.svcCtx.InfoRpc.DeleteEvaluation(l.ctx, &pb.DeleteEvaluationRequest{
		Pid: req.PostId,
	})
	if err != nil {
		resp.Code = 500
		resp.Msg = response.String()
	}
	return
}
