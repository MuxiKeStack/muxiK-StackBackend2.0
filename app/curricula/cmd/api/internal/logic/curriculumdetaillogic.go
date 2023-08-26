package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurriculumDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCurriculumDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurriculumDetailLogic {
	return &CurriculumDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CurriculumDetailLogic) CurriculumDetail(req *types.CurriculumDetailRequest) (resp *types.CurriculumDetailResponse, err error) {
	// todo: Test this method
	result, err := l.svcCtx.CurriculaCenterRpc.CurriculumDetail(l.ctx, &pb.CurriculumDetailRequest{
		DataId: req.DataId,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.CurriculumDetailResponse{
		CurriculaModel: types.CurriculaModel{
			DataId:          result.Info[0].DataId,
			CurriculaId:     result.Info[0].CurriculaId,
			CurriculaName:   result.Info[0].CurriculaName,
			Teacher:         result.Info[0].Teacher,
			Type:            uint8(result.Info[0].Type),
			Rate:            result.Info[0].Rate,
			StartsNum:       result.Info[0].StarsNum,
			GradeSampleSize: result.Info[0].GradeSampleSize,
			TotalGrade:      result.Info[0].TotalGrade,
			UsualGrade:      result.Info[0].UsualGrade,
			GradeRank1:      result.Info[0].GradeRk1,
			GradeRank2:      result.Info[0].GradeRk2,
			GradeRank3:      result.Info[0].GradeRk3,
		},
	}
	return resp, nil
}
