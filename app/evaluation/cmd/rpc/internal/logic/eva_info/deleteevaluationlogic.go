package evainfologic

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEvaluationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEvaluationLogic {
	return &DeleteEvaluationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteEvaluationLogic) DeleteEvaluation(in *pb.DeleteEvaluationRequest) (*pb.StatusResponse, error) {
	// todo: testing
	pid, err := strconv.Atoi(in.Pid)
	evaluationInfo, err := l.svcCtx.InfoModel.FindOne(l.ctx, int64(pid))
	evaluationInfo.Deleted = 1
	evaluationInfo.DeletedAt = sql.NullTime{Time: time.Now()}
	err = l.svcCtx.InfoModel.UpdateEvaluationInfo(l.ctx, evaluationInfo)
	if err != nil {
		return nil, err
	}
	return &pb.StatusResponse{Status: true}, nil
}
