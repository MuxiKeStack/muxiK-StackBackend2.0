package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpposeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpposeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpposeLogic {
	return &OpposeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OpposeLogic) Oppose(req *types.OpposeRequest) (resp *types.OpposeResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
