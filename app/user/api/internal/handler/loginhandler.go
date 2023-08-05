package handler

import (
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/api/internal/logic"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/api/internal/types"
	"github.com/Wishforpeace/zero-tools/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(w, resp, err)

	}
}
