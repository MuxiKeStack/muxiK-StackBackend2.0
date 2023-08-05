package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ccnu_one_login"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	if err := ccnu_one_login.LoginRequest(in.StudentID, in.Password); err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Token: "1111",
		IfNew: true,
	}, nil
}
