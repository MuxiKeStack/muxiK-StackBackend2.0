package message

import (
	"net/http"

	"github.com/big-dust/tool/result"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/cmd/api/internal/logic/message"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SocketConnectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ConnectReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := message.NewSocketConnectLogic(r.Context(), svcCtx)
		err := l.SocketConnect(w, r, &req)
		result.HttpResult(r, w, nil, err)
	}
}
