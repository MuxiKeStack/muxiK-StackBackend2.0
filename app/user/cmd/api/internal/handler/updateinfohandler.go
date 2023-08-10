package handler

import (
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/logic"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/types"
	"github.com/Wishforpeace/zero-tools/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func updateInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateInfoLogic(r.Context(), svcCtx)
		resp, err := l.UpdateInfo(&req)
		response.Response(w, resp, err)

	}
}
