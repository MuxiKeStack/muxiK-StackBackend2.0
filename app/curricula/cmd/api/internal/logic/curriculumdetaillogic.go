package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurriculumDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCurriculumDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurriculumDetailLogic {
	return &CurriculumDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CurriculumDetailLogic) CurriculumDetail(req *types.CurriculumDetailRequest) (resp *types.CurriculumDetailResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
