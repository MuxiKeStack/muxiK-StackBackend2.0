package message

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ctxdata"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/xerr"
	"github.com/jinzhu/copier"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOfficialMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOfficialMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOfficialMessageLogic {
	return &GetOfficialMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOfficialMessageLogic) GetOfficialMessage(req *types.GetOfficialMessageReq) (resp *types.GetOfficialMessageRsp, err error) {
	//获取群发
	publicMsgs, err := l.svcCtx.OfficialMessageModel.FindByNullSid(l.ctx)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ErrGetOfficialMessageByNullSid)
	}

	//获取单发
	privateMsgs, err := l.svcCtx.OfficialMessageModel.FindBySid(l.ctx, ctxdata.GetStudentIdFromCtx(l.ctx))
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ErrGetOfficialMessageBySid)
	}

	Msgs := append(publicMsgs, privateMsgs...)
	err = copier.Copy(&resp.OfficialMsgs, &Msgs)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ErrCopyData)
	}

	return
}
