package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckCharacteristicsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckCharacteristicsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckCharacteristicsLogic {
	return &CheckCharacteristicsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckCharacteristicsLogic) CheckCharacteristics(in *pb.CheckCharacteristicsRequest) (*pb.CheckCharacteristicsResponse, error) {
	// todo: Test this method
	var resp *pb.CheckCharacteristicsResponse
	result, err := l.svcCtx.CurriculaModel.FindByTag(l.ctx, uint8(in.Type))
	if err != nil {
		return nil, err
	}
	i := 0
	for result[i] != nil {
		resp.Info[i] = &pb.CurriculaInfo{
			DataId:        result[i].Id,
			CurriculaId:   uint32(result[i].Cid),
			CurriculaName: result[i].CurriculaName,
			Teacher:       result[i].Teacher,
			Type:          uint32(result[i].Type),
			Rate:          float32(result[i].Rate.Float64),
		}
		i++
	}
	return resp, nil
}
