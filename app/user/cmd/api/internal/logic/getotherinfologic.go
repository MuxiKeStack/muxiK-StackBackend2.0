package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/types"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOtherInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOtherInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOtherInfoLogic {
	return &GetOtherInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOtherInfoLogic) GetOtherInfo(req *types.OtherInfoRequest) (resp *types.OtherInfoResponse, err error) {

	infoResp, err := l.svcCtx.UserCenterRpc.GetInfo(l.ctx, &usercenter.GetInfoRequest{
		SID: req.StudentID,
	})
	if err != nil {
		return nil, err
	}
	return &types.OtherInfoResponse{
		OtherInfo: types.OtherInfo{
			StudentId: infoResp.StudentID,
			UserName:  infoResp.UserName,
			Avatar:    infoResp.Avatar,
		},
	}, nil
}
