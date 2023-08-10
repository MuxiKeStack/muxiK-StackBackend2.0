package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/pb"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/usercenter"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/model"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ccnu_one_login"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/helpers"
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
	var resp = new(pb.LoginResponse)
	if err := ccnu_one_login.LoginRequest(in.StudentID, in.Password); err != nil {
		return nil, err
	}
	var user *model.User
	//var err error
	num := l.svcCtx.UserModel.IfExist(l.ctx, in.StudentID)
	if num == 0 {
		user = &model.User{
			Sid:       in.StudentID,
			Username:  l.svcCtx.Config.Default.Username + helpers.GenerateRandomString(10),
			Avatar:    l.svcCtx.Config.Default.Avatar,
			IsBlocked: 0,
			RoleGrade: 0,
			Integral:  0,
			Licence:   0,
		}
		if err := l.svcCtx.UserModel.Insert(l.ctx, nil, user); err != nil {
			return nil, err
		}
		resp.IfNew = true
	} else {
		resp.IfNew = false
	}
	u, err := l.svcCtx.UserModel.FindOneBySid(l.ctx, in.StudentID)
	if err != nil {
		return nil, err
	}
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenRequest{
		StudentID: in.StudentID,
		RoleGrade: u.RoleGrade,
		Licence:   u.Licence,
	})
	if err != nil {
		return nil, err
	}
	resp.Token = tokenResp.Token
	return resp, nil
}
