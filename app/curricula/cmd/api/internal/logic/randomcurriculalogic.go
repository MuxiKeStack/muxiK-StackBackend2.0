package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RandomCurriculaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRandomCurriculaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RandomCurriculaLogic {
	return &RandomCurriculaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RandomCurriculaLogic) RandomCurricula(req *types.RandomCurriculaRequest) (resp *types.RandomCurriculaResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
