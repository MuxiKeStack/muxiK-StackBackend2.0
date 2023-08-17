package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/model"
	rds "github.com/MuxiKeStack/muxiK-StackBackend2.0/common/redis"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/redis_store"
	"strconv"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseIntegralLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncreaseIntegralLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseIntegralLogic {
	return &IncreaseIntegralLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IncreaseIntegralLogic) IncreaseIntegral(in *pb.IncreaseIntegralRequest) (*pb.IncreaseIntegralResponse, error) {
	var user *model.User
	var err error
	if in.Type == "like" {
		if l.overIntegral("like_integral_", in.StudentId, 1) {
			user, err = l.svcCtx.UserModel.IncreaseIntegral(l.ctx, in.StudentId, in.Integral)
			if err != nil {
				return nil, err
			}
		} else {
			user, err = l.svcCtx.UserModel.FindOneBySid(l.ctx, in.StudentId)
			if err != nil {
				return nil, err
			}
		}
	}
	if in.Type == "comment" {
		if l.overIntegral("comment_integral_", in.StudentId, 4) {
			user, err = l.svcCtx.UserModel.IncreaseIntegral(l.ctx, in.StudentId, in.Integral)
			if err != nil {
				return nil, err
			}
		} else {
			user, err = l.svcCtx.UserModel.FindOneBySid(l.ctx, in.StudentId)
			return nil, err
		}
	}
	if in.Type != "like" && in.Type != "comment" {
		user, err = l.svcCtx.UserModel.IncreaseIntegral(l.ctx, in.StudentId, in.Integral)
		if err != nil {
			return nil, err
		}
	}

	return &pb.IncreaseIntegralResponse{
		StudentId: user.Sid,
		Integral:  user.Integral,
		RoleGrade: user.RoleGrade,
	}, nil
}

func (l *IncreaseIntegralLogic) overIntegral(prefix string, sid string, limit int) bool {
	var integral = redis_store.RedisStore{
		RedisClient: &rds.RedisClient{
			Client:  l.svcCtx.RedisClient,
			Context: l.ctx,
		},
		KeyPrefix: prefix,
	}
	result, _ := strconv.Atoi(integral.Get(sid, false))
	if result >= limit {
		return false
	}
	return true
}
