package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCurriculaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCurriculaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCurriculaLogic {
	return &DeleteCurriculaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCurriculaLogic) DeleteCurricula(in *pb.DeleteCurriculaRequest) (*pb.DeleteCurriculaResponse, error) {
	// todo: Test this method
	err := l.svcCtx.CurriculaModel.Delete(l.ctx, in.DataId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCurriculaResponse{}, nil
}
