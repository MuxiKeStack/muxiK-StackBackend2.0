package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/types"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/curriculacenter"

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
	_, err = l.svcCtx.CurriculaCenterRpc.AddCurricula(l.ctx, &curriculacenter.AddCurriculaRequest{
		CurriculaName: req.CurriculaNmae,
		CurriculaId:   req.CurriculaId,
		Type:          req.Type,
		Teacher:       req.Teacher,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddCurriculaResponse{}, nil
}
