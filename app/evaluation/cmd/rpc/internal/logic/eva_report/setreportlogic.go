package evareportlogic

import (
	"context"
	"database/sql"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/model"
	"strconv"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetReportLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetReportLogic {
	return &SetReportLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetReportLogic) SetReport(in *pb.SetReportRequest) (*pb.StatusResponse, error) {
	// todo: test
	i, err := strconv.Atoi(in.R.Status)
	err = l.svcCtx.ReportModel.Update(l.ctx, &model.EvaluationReport{
		Pid: in.R.Pid,
		Sid: in.R.Sid,
		Reason: sql.NullString{
			String: in.R.Reason,
			Valid:  true,
		},

		Status: sql.NullInt64{
			Int64: int64(i),
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.StatusResponse{Status: true}, nil
}
