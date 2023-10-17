package logic

import (
	"context"
	"fmt"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/model"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ctxdata"

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
	// todo: add your logic here and delete this line
	post := model.EvaluationInfo{
		Sid:  ctxdata.GetStudentIdFromCtx(l.ctx),
		Cid:  req.CourseId,
		Info: req.Info,
	}
	fmt.Println(post)
	return
}
