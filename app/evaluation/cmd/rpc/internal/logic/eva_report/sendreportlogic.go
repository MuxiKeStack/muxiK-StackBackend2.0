package evareportlogic

import (
	"context"
	"database/sql"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/model"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendReportLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendReportLogic {
	return &SendReportLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendReportLogic) SendReport(in *pb.SendReportRequest) (*pb.StatusResponse, error) {
	// todo: test
	_, err := l.svcCtx.ReportModel.Insert(l.ctx, &model.EvaluationReport{
		Pid: in.R.Pid,
		Sid: in.R.Sid,
		Reason: sql.NullString{
			String: in.R.Reason,
			Valid:  true,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.StatusResponse{Status: true}, nil
}
