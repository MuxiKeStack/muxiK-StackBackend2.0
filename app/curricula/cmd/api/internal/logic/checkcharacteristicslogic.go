package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"

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
	// todo: Test this method
	result, err := l.svcCtx.CurriculaCenterRpc.CheckCharacteristics(l.ctx, &pb.CheckCharacteristicsRequest{Type: uint32(req.Type)})
	if err != nil {
		return nil, err
	}
	i := 0
	for result.Info[i] != nil {
		resp.Info[i] = types.CurriculaInfo{
			DataId:        result.Info[i].DataId,
			CurriculaId:   result.Info[i].CurriculaId,
			CurriculaName: result.Info[i].CurriculaName,
			Teacher:       result.Info[i].Teacher,
			Type:          result.Info[i].Type,
			Rate:          result.Info[i].Rate,
		}
	}
	return resp, nil
}
