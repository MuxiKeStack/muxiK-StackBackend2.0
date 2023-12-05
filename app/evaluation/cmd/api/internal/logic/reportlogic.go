package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/types"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ctxdata"

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
	// todo: test

	_, err = l.svcCtx.ReportRpc.SendReport(l.ctx, &pb.SendReportRequest{
		R: &pb.Report{
			Pid:    req.PostId,
			Sid:    ctxdata.GetStudentIdFromCtx(l.ctx),
			Reason: req.Reason,
			Status: "0",
		}},
	)
	if err != nil {
		return nil, err
	}

	return
}
