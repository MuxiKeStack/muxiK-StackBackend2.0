package reply

import (
	"net/http"

	"github.com/big-dust/tool/result"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation1/cmd/api/internal/logic/reply"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation1/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation1/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReplyCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReplyCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := reply.NewReplyCommentLogic(r.Context(), svcCtx)
		resp, err := l.ReplyComment(&req)
		result.HttpResult(r, w, resp, err)
	}
}
