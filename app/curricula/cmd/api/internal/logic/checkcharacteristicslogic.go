package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckCharacteristicsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckCharacteristicsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckCharacteristicsLogic {
	return &CheckCharacteristicsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckCharacteristicsLogic) CheckCharacteristics(req *types.CheckCharacteristicsRequest) (resp *types.CheckCharacteristicsResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
