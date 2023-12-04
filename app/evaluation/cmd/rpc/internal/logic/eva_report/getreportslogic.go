package evareportlogic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/model"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetReportsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetReportsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReportsLogic {
	return &GetReportsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetReportsLogic) GetReports(in *pb.GetReportsRequest) (*pb.GetReportsResponse, error) {
	// todo: test
	data, err := l.svcCtx.ReportModel.FindAll(l.ctx)

	if err != nil {
		return nil, err
	}
	return &pb.GetReportsResponse{Reports: func(data []*model.EvaluationReport) []*pb.Report {
		ret := make([]*pb.Report, len(data))
		for i, v := range data {
			ret[i] = &pb.Report{
				Pid:    v.Pid,
				Sid:    v.Sid,
				Reason: v.Reason.String,
			}
		}
		return ret
	}(data)}, nil
}
