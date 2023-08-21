package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCurriculaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCurriculaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCurriculaLogic {
	return &DeleteCurriculaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCurriculaLogic) DeleteCurricula(req *types.DeleteCurriculaRequest) (resp *types.DeleteCurriculaResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
