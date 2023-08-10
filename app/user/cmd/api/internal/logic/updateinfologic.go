package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/usercenter"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ctxdata"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInfoLogic {
	return &UpdateInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateInfoLogic) UpdateInfo(req *types.UpdateInfoRequest) (resp *types.UpdateInfoResponse, err error) {
	sid := ctxdata.GetStudentIdFromCtx(l.ctx)
	_, err = l.svcCtx.UserCenterRpc.UpdateInfo(l.ctx, &usercenter.UpdateInfoRequest{
		SID:      sid,
		Avatar:   req.Avatar,
		UserName: req.UserName,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateInfoResponse{}, nil
}
