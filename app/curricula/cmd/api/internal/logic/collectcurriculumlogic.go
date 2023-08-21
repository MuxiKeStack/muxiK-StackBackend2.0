package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectCurriculumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectCurriculumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectCurriculumLogic {
	return &CollectCurriculumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectCurriculumLogic) CollectCurriculum(req *types.CollectCurriculumRequest) (resp *types.CollectCurriculumResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
