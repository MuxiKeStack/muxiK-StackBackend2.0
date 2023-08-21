package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCurriculaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCurriculaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCurriculaLogic {
	return &UpdateCurriculaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCurriculaLogic) UpdateCurricula(in *pb.UpdateCurriculaRequest) (*pb.UpdateCurriculaResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateCurriculaResponse{}, nil
}
