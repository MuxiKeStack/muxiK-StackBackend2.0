package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &pb.StatusResponse{}, nil
}
