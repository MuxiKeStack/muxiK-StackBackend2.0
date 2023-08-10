package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/types"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/usercenter"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	loginResp, err := l.svcCtx.UserCenterRpc.Login(l.ctx, &usercenter.LoginRequest{
		StudentID: req.StudentId,
		Password:  req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.LoginResponse{
		Token: loginResp.Token,
		IfNew: loginResp.IfNew,
	}, nil
}
