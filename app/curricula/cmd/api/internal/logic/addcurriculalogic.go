package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCurriculaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCurriculaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCurriculaLogic {
	return &AddCurriculaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCurriculaLogic) AddCurricula(req *types.AddCurriculaRequest) (resp *types.AddCurriculaResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
