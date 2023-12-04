package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"

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
	// todo: test this method
	response, err := l.svcCtx.CurriculaCenterRpc.RandomCurricula(l.ctx, &pb.RandomRequest{})
	if err != nil {
		return nil, err
	}
	for _, info := range response.Info {
		resp.Info = append(resp.Info, types.CurriculaInfo{
			DataId:        info.DataId,
			CurriculaId:   info.CurriculaId,
			CurriculaName: info.CurriculaName,
			Teacher:       info.Teacher,
			Type:          info.Type,
			Rate:          info.Rate,
		})
	}
	return
}
