package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/usercenter"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ctxdata"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SharingPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSharingPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SharingPlanLogic {
	return &SharingPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SharingPlanLogic) SharingPlan(req *types.SharingRequest) (resp *types.SharingResponse, err error) {
	_, err = l.svcCtx.UserCenterRpc.SharingPlan(l.ctx, &usercenter.SharingRequest{
		SID: ctxdata.GetStudentIdFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}
	return &types.SharingResponse{}, nil

}
