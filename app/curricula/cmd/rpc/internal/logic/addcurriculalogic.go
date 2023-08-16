package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCurriculaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCurriculaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCurriculaLogic {
	return &AddCurriculaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCurriculaLogic) AddCurricula(in *pb.AddCurriculaRequest) (*pb.AddCurriculaResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.AddCurriculaResponse{}, nil
}
