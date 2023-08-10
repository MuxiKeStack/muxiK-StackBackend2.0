package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoLogic {
	return &GetInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInfoLogic) GetInfo(in *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	user, err := l.svcCtx.UserModel.FindOneBySid(l.ctx, in.SID)
	if err != nil {
		return nil, err
	}
	return &pb.GetInfoResponse{
		StudentID: user.Sid,
		UserName:  user.Username,
		Avatar:    user.Avatar,
		IsBlocked: user.IsBlocked,
		RoleGrade: user.RoleGrade,
		Integral:  user.Integral,
		Licence:   user.Licence,
	}, nil
}
