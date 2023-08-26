package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurriculumDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCurriculumDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurriculumDetailLogic {
	return &CurriculumDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CurriculumDetailLogic) CurriculumDetail(in *pb.CurriculumDetailRequest) (*pb.CurriculumDetailResponse, error) {
	// todo: Test this method
	curricula, err := l.svcCtx.CurriculaModel.FindOne(l.ctx, in.DataId)
	if err != nil {
		return nil, err
	}
	var res pb.CurriculumDetailResponse
	res.Info[0] = &pb.CurriculaModel{
		DataId:          curricula.Id,
		CurriculaId:     uint32(curricula.Cid),
		CurriculaName:   curricula.CurriculaName,
		Teacher:         curricula.Teacher,
		Type:            uint32(curricula.Type),
		Rate:            float32(curricula.Rate.Float64),
		StarsNum:        uint32(curricula.StartsNum.Int64),
		GradeSampleSize: uint32(curricula.GradeSampleSize.Int64),
		TotalGrade:      float32(curricula.TotalGrade.Float64),
		UsualGrade:      float32(curricula.UsualGrade.Float64),
		GradeRk1:        uint32(curricula.GradeR1.Int64),
		GradeRk2:        uint32(curricula.GradeR2.Int64),
		GradeRk3:        uint32(curricula.GradeR3.Int64),
	}
	return &res, nil
}
