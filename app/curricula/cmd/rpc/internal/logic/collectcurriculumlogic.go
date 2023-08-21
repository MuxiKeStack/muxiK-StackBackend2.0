package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectCurriculumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectCurriculumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectCurriculumLogic {
	return &CollectCurriculumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CollectCurriculumLogic) CollectCurriculum(in *pb.CollectCurriculumRequest) (*pb.CollectCurriculumResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CollectCurriculumResponse{}, nil
}
