package evainfologic

import (
	"context"
	"database/sql"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/model"
	"time"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEvaluationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEvaluationLogic {
	return &CreateEvaluationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateEvaluationLogic) CreateEvaluation(in *pb.CreateEvaluationRequest) (*pb.StatusResponse, error) {
	// todo: testing
	err := l.svcCtx.InfoModel.CreateEvaluationInfo(l.ctx, &model.EvaluationInfo{
		Sid:       in.E.Sid,
		Cid:       in.E.Cid,
		Folded:    0,
		Deleted:   0,
		Info:      in.E.Info,
		Liked:     sql.NullInt64{},
		Disliked:  sql.NullInt64{},
		CreatedAt: sql.NullTime{Time: time.Now()},
		UpdatedAt: sql.NullTime{Time: time.Now()},
		DeletedAt: sql.NullTime{},
	})
	if err != nil {
		return &pb.StatusResponse{Status: false}, err
	}
	return &pb.StatusResponse{Status: true}, nil
}
