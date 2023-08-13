package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/pb"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/usercenter"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/model"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ccnu_one_login"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/helpers"
	rds "github.com/MuxiKeStack/muxiK-StackBackend2.0/common/redis"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/redis_store"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
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
	err = l.setIntegral(in.StudentID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LoginLogic) setIntegral(sid string) error {
	tomorrow := time.Now().Add(24 * time.Hour)
	expiration := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, time.UTC)
	duration := expiration.Sub(time.Now())
	var likeIntegral = redis_store.RedisStore{
		RedisClient: &rds.RedisClient{
			Client:  l.svcCtx.RedisClient,
			Context: l.ctx,
		},
		KeyPrefix: "like_integral_",
	}
	likeResult := likeIntegral.Get(sid, false)
	if likeResult == "" {
		if err := likeIntegral.Set(sid, 0, duration); err != nil {
			return err
		}
	}
	var commentIntegral = redis_store.RedisStore{
		RedisClient: &rds.RedisClient{
			Client:  l.svcCtx.RedisClient,
			Context: l.ctx,
		},
		KeyPrefix: "comment_integral_",
	}
	commentResult := commentIntegral.Get(sid, false)
	if commentResult == "" {
		if err := commentIntegral.Set(sid, 0, duration); err != nil {
			return err
		}
	}
	return nil
}
