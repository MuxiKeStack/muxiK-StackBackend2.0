package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckCharacteristicsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckCharacteristicsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckCharacteristicsLogic {
	return &CheckCharacteristicsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckCharacteristicsLogic) CheckCharacteristics(in *pb.CheckCharacteristicsRequest) (*pb.CheckCharacteristicsResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CheckCharacteristicsResponse{}, nil
}
