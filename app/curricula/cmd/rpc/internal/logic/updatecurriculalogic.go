package logic

import (
	"context"
	"database/sql"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/model"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/xerr"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCurriculaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCurriculaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCurriculaLogic {
	return &UpdateCurriculaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCurriculaLogic) UpdateCurricula(in *pb.UpdateCurriculaRequest) (*pb.UpdateCurriculaResponse, error) {
	// todo: test this method
	curricula := in.Model[0]
	err := l.svcCtx.CurriculaModel.Update(l.ctx, &model.Curriculas{
		Id:              curricula.DataId,
		Cid:             int64(curricula.CurriculaId),
		CurriculaName:   curricula.CurriculaName,
		Teacher:         curricula.Teacher,
		Type:            int64(curricula.Type),
		Rate:            sql.NullFloat64{Float64: float64(curricula.Rate), Valid: true},
		StartsNum:       sql.NullInt64{Int64: int64(curricula.StarsNum), Valid: true},
		GradeSampleSize: sql.NullInt64{Int64: int64(curricula.GradeSampleSize), Valid: true},
		TotalGrade:      sql.NullFloat64{Float64: float64(curricula.TotalGrade), Valid: true},
		UsualGrade:      sql.NullFloat64{Float64: float64(curricula.UsualGrade), Valid: true},
		GradeR1:         sql.NullInt64{Int64: int64(curricula.GradeRk1), Valid: true},
		GradeR2:         sql.NullInt64{Int64: int64(curricula.GradeRk2), Valid: true},
		GradeR3:         sql.NullInt64{Int64: int64(curricula.GradeRk3), Valid: true},
	})

	if err != nil {
		return nil, xerr.NewErrCode(xerr.ErrUpdateCourseInfo)
	}
	return &pb.UpdateCurriculaResponse{}, nil
}
