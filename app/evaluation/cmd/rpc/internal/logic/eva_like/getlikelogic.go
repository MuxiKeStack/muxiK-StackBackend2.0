package evalikelogic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLikeLogic {
	return &GetLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLikeLogic) GetLike(in *pb.GetLikeRequest) (*pb.GetLikeResponse, error) {
	// todo: testing
	info, err := l.svcCtx.LikeModel.FindLike(l.ctx, in.Pid, in.Sid)
	if err != nil {
		return nil, err
	}
	return &pb.GetLikeResponse{L: &pb.Like{
		Pid:    in.Pid,
		Sid:    in.Sid,
		Status: int32(info.Status),
	}}, nil
}
