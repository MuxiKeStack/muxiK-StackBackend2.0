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

type GetNormalMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNormalMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNormalMessageLogic {
	return &GetNormalMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNormalMessageLogic) GetNormalMessage(req *types.GetNormalMessageReq) (resp *types.GetNormalMessageRsp, err error) {

	normalMessages, err := l.svcCtx.NormalMessageModel.FindBySid(l.ctx, ctxdata.GetStudentIdFromCtx(l.ctx))
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ErrGetNormalMessageBySid)
	}

	err = copier.Copy(&resp.NormalMsgs, &normalMessages)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ErrCopyData)
	}
	return
}
