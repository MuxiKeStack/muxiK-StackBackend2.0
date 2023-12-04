package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCurriculaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCurriculaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCurriculaLogic {
	return &UpdateCurriculaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCurriculaLogic) UpdateCurricula(req *types.UpdateCurriculaRequest) (resp *types.UpdateCurriculaResponse, err error) {
	// todo: test this method
	var models []*pb.CurriculaModel
	models[0] = &pb.CurriculaModel{
		DataId:          req.DataId,
		CurriculaId:     req.CurriculaId,
		CurriculaName:   req.CurriculaName,
		Teacher:         req.Teacher,
		Type:            uint32(req.Type),
		Rate:            req.Rate,
		StarsNum:        req.StartsNum,
		GradeSampleSize: req.GradeSampleSize,
		TotalGrade:      req.TotalGrade,
		UsualGrade:      req.UsualGrade,
		GradeRk1:        req.GradeRank1,
		GradeRk2:        req.GradeRank2,
		GradeRk3:        req.GradeRank3,
	}
	_, err = l.svcCtx.CurriculaCenterRpc.UpdateCurricula(l.ctx, &pb.UpdateCurriculaRequest{Model: models})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
