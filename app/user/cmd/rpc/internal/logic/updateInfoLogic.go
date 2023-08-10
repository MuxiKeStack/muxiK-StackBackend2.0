package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInfoLogic {
	return &UpdateInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateInfoLogic) UpdateInfo(in *pb.UpdateInfoRequest) (*pb.UpdateInfoResponse, error) {
	user, err := l.svcCtx.UserModel.FindOneBySid(l.ctx, in.SID)
	if err != nil {
		return nil, err
	}
	if in.Avatar != "" {
		user.Avatar = in.Avatar
	}
	if in.UserName != "" {
		user.Username = in.UserName
	}
	if err := l.svcCtx.UserModel.Update(l.ctx, nil, user); err != nil {
		return nil, err
	}
	return &pb.UpdateInfoResponse{}, nil
}
