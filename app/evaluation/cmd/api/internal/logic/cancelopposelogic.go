package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ctxdata"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelOpposeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelOpposeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelOpposeLogic {
	return &CancelOpposeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelOpposeLogic) CancelOppose(req *types.CancelOpposeRequest) (resp *types.CancelOpposeResponse, err error) {
	// todo: test
	l.svcCtx.LikeRpc.SetLike(l.ctx, &pb.SetLikeRequest{
		Pid:    req.PostId,
		Sid:    ctxdata.GetStudentIdFromCtx(l.ctx),
		Status: 0,
	})
	return
}
