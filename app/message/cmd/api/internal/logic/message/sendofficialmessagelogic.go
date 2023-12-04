package message

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/cmd/api/internal/types"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/model"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/xerr"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendOfficialMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendOfficialMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendOfficialMessageLogic {
	return &SendOfficialMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendOfficialMessageLogic) SendOfficialMessage(req *types.SendOfficialMessageReq) (resp *types.SendOfficialMessageRsp, err error) {

	raw, err := json.Marshal(req.Msg)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ErrMarshal)
	}

	if req.AllUser {
		err = l.svcCtx.Hub.Broadcast(raw)
		if err != nil {
			return nil, xerr.NewErrCode(xerr.ErrBroadcast)
		}
		req.Msg.ObjectSid = ""
	} else {
		err = l.svcCtx.Hub.Send(req.Msg.ObjectSid, raw)
		if err != nil {
			return nil, xerr.NewErrCode(xerr.ErrSendMessage)
		}
	}

	//存数据库
	_, err = l.svcCtx.OfficialMessageModel.Insert(l.ctx, &model.OfficialMessage{
		Title: req.Msg.Title,
		Text:  req.Msg.Text,
		Image: sql.NullString{
			String: req.Msg.Image,
			Valid:  req.Msg.Image != "",
		},
		ObjectSid: sql.NullString{
			String: req.Msg.ObjectSid,
			Valid:  req.Msg.ObjectSid != "",
		},
		SendAt: time.Now().Unix(),
	})

	if err != nil {
		return nil, xerr.NewErrCode(xerr.ErrStoreOfficialMessage)
	}
	resp.Flag = true
	return
}
