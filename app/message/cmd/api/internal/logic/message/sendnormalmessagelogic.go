package message

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/model"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ctxdata"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/xerr"
	"time"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendNormalMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendNormalMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendNormalMessageLogic {
	return &SendNormalMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendNormalMessageLogic) SendNormalMessage(req *types.SendNormalMessageReq) (resp *types.SendNormalMessageRsp, err error) {

	message := &types.NormalMessage{
		SenderId:        ctxdata.GetStudentIdFromCtx(l.ctx),
		Type:            req.Type,
		ObjectSid:       req.ObjectSid,
		OriginContentId: req.OriginContentId,
		Text:            req.Text,
		SendAt:          time.Now().Unix(),
	}

	raw, err := json.Marshal(message)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ErrMarshal)
	}

	err = l.svcCtx.Hub.Send(req.ObjectSid, raw)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ErrSendMessage)
	}

	// 存数据库
	_, err = l.svcCtx.NormalMessageModel.Insert(l.ctx, &model.NormalMessage{
		Type:            req.Type,
		SenderSid:       ctxdata.GetStudentIdFromCtx(l.ctx),
		ObjectSid:       req.ObjectSid,
		OriginContentId: int64(req.OriginContentId),
		Text: sql.NullString{
			String: req.Text,
			Valid:  req.Text != "",
		},
		SendAt: time.Now().Unix(),
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ErrStoreNormalMessage)
	}

	resp.Flag = true
	return
}
