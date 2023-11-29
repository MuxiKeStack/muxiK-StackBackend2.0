package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &pb.GetReportsResponse{}, nil
}
