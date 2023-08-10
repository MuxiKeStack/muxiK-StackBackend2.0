package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SharingPlanLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSharingPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SharingPlanLogic {
	return &SharingPlanLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SharingPlanLogic) SharingPlan(in *pb.SharingRequest) (*pb.SharingResponse, error) {
	user, err := l.svcCtx.UserModel.FindOneBySid(l.ctx, in.SID)
	if err != nil {
		return nil, err
	}
	user.Licence = 1
	if err := l.svcCtx.UserModel.Update(l.ctx, nil, user); err != nil {
		return nil, err
	}

	return &pb.SharingResponse{}, nil
}
