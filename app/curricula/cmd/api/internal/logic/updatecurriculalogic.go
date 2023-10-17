package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCurriculaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCurriculaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCurriculaLogic {
	return &UpdateCurriculaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCurriculaLogic) UpdateCurricula(req *types.UpdateCurriculaRequest) (resp *types.UpdateCurriculaResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
