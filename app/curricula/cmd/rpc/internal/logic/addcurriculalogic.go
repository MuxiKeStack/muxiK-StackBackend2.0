package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/model"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/xerr"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddCurriculaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCurriculaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCurriculaLogic {
	return &AddCurriculaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCurriculaLogic) AddCurricula(in *pb.AddCurriculaRequest) (*pb.AddCurriculaResponse, error) {
	curricula, err := l.svcCtx.CurriculaModel.FindOneByCid(l.ctx, int64(in.CurriculaId))
	if err == nil {
		return nil, xerr.NewErrCode(xerr.ErrCourseHasExisted)
	}
	curricula = &model.Curriculas{
		Cid:           int64(in.CurriculaId),
		CurriculaName: in.CurriculaName,
		Teacher:       in.Teacher,
		Type:          int64(in.Type),
	}
	_, err = l.svcCtx.CurriculaModel.Insert(l.ctx, curricula)
	if err != nil {
		return nil, err
	}
	return &pb.AddCurriculaResponse{}, nil
}
