package evareportlogic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTheReportLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTheReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTheReportLogic {
	return &GetTheReportLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTheReportLogic) GetTheReport(in *pb.GetTheReportRequest) (*pb.GetTheReportResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetTheReportResponse{}, nil
}
