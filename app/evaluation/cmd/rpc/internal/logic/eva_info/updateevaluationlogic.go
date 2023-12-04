package evainfologic

import (
	"context"
	"database/sql"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/model"
	"strconv"
	"time"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEvaluationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEvaluationLogic {
	return &UpdateEvaluationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func atoi(s string) (a int) {
	a, _ = strconv.Atoi(s)
	return a
}

func btoi(b bool) (a int) {
	if b {
		return 1
	}
	return 0
}

func gettime(s string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", s)
	return t
}

func (l *UpdateEvaluationLogic) UpdateEvaluation(in *pb.UpdateEvaluationRequest) (*pb.StatusResponse, error) {
	// todo: testing
	err := l.svcCtx.InfoModel.UpdateEvaluationInfo(l.ctx, &model.EvaluationInfo{
		Pid:     int64(atoi(in.E.Pid)),
		Sid:     in.E.Sid,
		Cid:     in.E.Cid,
		Folded:  int64(btoi(in.E.Folded)),
		Deleted: int64(btoi(in.E.Deleted)),
		Info:    in.E.Info,
		Liked: sql.NullInt64{
			Int64: int64(in.E.Liked),
			Valid: true,
		},
		Disliked: sql.NullInt64{
			Int64: int64(in.E.Disliked),
			Valid: true,
		},
		CreatedAt: sql.NullTime{
			Time: gettime(in.E.CreateAt),
		},
		UpdatedAt: sql.NullTime{
			Time: gettime(in.E.UpdateAt),
		},
		DeletedAt: sql.NullTime{
			Time: gettime(in.E.DeleteAt),
		},
	})
	if err != nil {
		return &pb.StatusResponse{Status: false}, err
	}
	return &pb.StatusResponse{Status: true}, nil
}
