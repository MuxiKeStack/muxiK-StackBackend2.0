package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/types"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHistoryLogic {
	return &GetHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHistoryLogic) GetHistory(req *types.GetHistoryRequest) (resp *types.GetHistoryResponse, err error) {
	// todo: add your logic here and delete this line
	evaluation, err := l.svcCtx.InfoRpc.GetEvaluation(l.ctx, &pb.GetEvaluationRequest{
		Method: "sid",
		Id:     req.StudentId,
	})
	if err != nil {
		return nil, err
	}

	for _, v := range evaluation.E {
		resp.Info = append(resp.Info, types.EvaluationInfo{
			PostId:    v.Pid,
			StudentId: v.Sid,
			CourseId:  v.Cid,
			Info:      v.Info,
		})
	}

	return
}
