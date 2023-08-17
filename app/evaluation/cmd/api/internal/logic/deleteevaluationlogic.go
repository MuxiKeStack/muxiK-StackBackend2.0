package logic

import (
	"context"
	"fmt"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEvaluationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEvaluationLogic {
	return &DeleteEvaluationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteEvaluationLogic) DeleteEvaluation(req *types.DeleteEvaluationRequest) (resp *types.DeleteEvaluationResponse, err error) {
	fmt.Println(l.ctx.Value("pid"))
	req.PostId = l.ctx.Value("pid").(string)
	// todo: add your logic here and delete this line

	return
}
