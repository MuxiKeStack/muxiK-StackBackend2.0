package comment

import (
	"net/http"

	"github.com/big-dust/tool/result"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation(1)/cmd/api/internal/logic/comment"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation(1)/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation(1)/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CommentEvaluationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentEvaluationReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := comment.NewCommentEvaluationLogic(r.Context(), svcCtx)
		resp, err := l.CommentEvaluation(&req)
		result.HttpResult(r, w, resp, err)
	}
}
