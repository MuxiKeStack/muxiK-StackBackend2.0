package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/types"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeLogic) Like(req *types.LikeRequest) (resp *types.LikeResponse, err error) {
	// todo: test
	l.svcCtx.LikeRpc.SetLike(l.ctx, &pb.SetLikeRequest{
		Pid:    req.PostId,
		Sid:    ctxdata.GetStudentIdFromCtx(l.ctx),
		Status: 1,
	})
	return
}
