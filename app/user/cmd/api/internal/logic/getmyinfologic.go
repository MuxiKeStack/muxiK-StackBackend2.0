package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/usercenter"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ctxdata"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyInfoLogic {
	return &GetMyInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyInfoLogic) GetMyInfo(req *types.MyInfoRequest) (resp *types.MyInfoResponse, err error) {
	sid := ctxdata.GetStudentIdFromCtx(l.ctx)
	infoResp, err := l.svcCtx.UserCenterRpc.GetInfo(l.ctx, &usercenter.GetInfoRequest{
		SID: sid,
	})
	if err != nil {
		return nil, err
	}
	return &types.MyInfoResponse{
		PersonalInfo: types.PersonalInfo{
			StudentId: infoResp.StudentID,
			UserName:  infoResp.UserName,
			Avatar:    infoResp.Avatar,
			IsBlocked: infoResp.IsBlocked,
			RoleGrade: infoResp.RoleGrade,
			Integral:  infoResp.Integral,
			Licence:   infoResp.Licence,
		},
	}, nil
}
