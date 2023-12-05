package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ctxdata"
	"time"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EvaluationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EvaluationLogic {
	return &EvaluationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EvaluationLogic) Evaluation(req *types.EvaluateRequest) (resp *types.EvaluateResponse, err error) {
	// todo: test
	r, err := l.svcCtx.InfoRpc.CreateEvaluation(l.ctx, &pb.CreateEvaluationRequest{E: &pb.Evaluation{
		Sid:      ctxdata.GetStudentIdFromCtx(l.ctx),
		Cid:      req.CourseId,
		Info:     req.Info,
		Liked:    0,
		Disliked: 0,
		CreateAt: time.Now().String(),
	}})
	if r.Status == false {
		resp.Code = 400
		resp.Msg = err.Error()
	}
	return
}
