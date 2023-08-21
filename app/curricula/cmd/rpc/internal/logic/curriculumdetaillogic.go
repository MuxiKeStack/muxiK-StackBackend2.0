package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurriculumDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCurriculumDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurriculumDetailLogic {
	return &CurriculumDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CurriculumDetailLogic) CurriculumDetail(in *pb.CurriculumDetailRequest) (*pb.CurriculumDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CurriculumDetailResponse{}, nil
}
