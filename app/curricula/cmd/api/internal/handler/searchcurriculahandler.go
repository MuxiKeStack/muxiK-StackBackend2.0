package handler

import (
	"net/http"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/logic"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func searchCurriculaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchCurriculaRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSearchCurriculaLogic(r.Context(), svcCtx)
		resp, err := l.SearchCurricula(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}