package handler

import (
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/qiniu/cmd/api/internal/logic/qiniu"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/qiniu/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/qiniu/cmd/api/internal/types"
	"github.com/Wishforpeace/zero-tools/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func Qiniu_tokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QiniuTokenRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := qiniu.NewQiniu_tokenLogic(r.Context(), svcCtx)
		resp, err := l.Qiniu_token(&req)
		response.Response(w, resp, err)

	}
}
