package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/model"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCurriculaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchCurriculaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCurriculaLogic {
	return &SearchCurriculaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchCurriculaLogic) SearchCurricula(in *pb.SearchCurriculaRequest) (*pb.SearchCurriculaResponse, error) {
	var curricula model.Cinfo
	curricula = model.Cinfo{
		CurriculaId:   in.CurriculaId,
		CurriculaName: in.CurriculaName,
		Teacher:       in.Teacher,
		Type:          uint8(in.Type),
	}
	resp, err := l.svcCtx.CurriculaModel.FindByInfos(l.ctx, curricula)
	if err != nil {
		return &pb.SearchCurriculaResponse{}, err
	}
	var res *pb.SearchCurriculaResponse
	res.Info[0] = &pb.CurriculaInfo{
		DataId:        resp.Id,
		CurriculaId:   uint32(resp.Cid),
		CurriculaName: resp.CurriculaName,
		Teacher:       resp.Teacher,
		Type:          uint32(resp.Type),
		Rate:          float32(resp.Rate.Float64),
	}
	return res, nil
}
