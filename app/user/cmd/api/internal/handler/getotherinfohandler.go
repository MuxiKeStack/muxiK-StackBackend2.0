package handler

import (
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/logic"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/types"
	"github.com/Wishforpeace/zero-tools/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func getOtherInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OtherInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetOtherInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetOtherInfo(&req)
		response.Response(w, resp, err)

	}
}
