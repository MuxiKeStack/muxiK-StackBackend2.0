package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelOpposeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelOpposeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelOpposeLogic {
	return &CancelOpposeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelOpposeLogic) CancelOppose(req *types.CancelOpposeRequest) (resp *types.CancelOpposeResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
