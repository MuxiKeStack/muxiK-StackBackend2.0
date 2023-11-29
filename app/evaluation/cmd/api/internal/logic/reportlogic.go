package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReportLogic {
	return &ReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReportLogic) Report(req *types.ReportRequest) (resp *types.ReportResponse, err error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.InfoRpc.UpdateEvaluation(l.ctx, &pb.UpdateEvaluationRequest{E: &pb.Evaluation{
		Pid: req.PostId,
	}})
	if err != nil {
		resp.Code = 500
		resp.Msg = response.String()
	}
	return
}
