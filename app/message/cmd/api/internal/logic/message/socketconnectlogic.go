package message

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/cmd/api/internal/socket"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ctxdata"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/xerr"
	"net/http"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SocketConnectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSocketConnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SocketConnectLogic {
	return &SocketConnectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SocketConnectLogic) SocketConnect(w http.ResponseWriter, r *http.Request, req *types.ConnectReq) error {
	sid := ctxdata.GetStudentIdFromCtx(l.ctx)
	err := socket.ServeWs(sid, l.svcCtx.Hub, w, r)
	if err != nil {
		return xerr.NewErrCode(xerr.ErrUpgradeWebsocket)
	}
	return nil
}
