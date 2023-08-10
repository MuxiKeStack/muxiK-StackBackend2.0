package middleware

import (
	"fmt"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ctxdata"
	"net/http"
)

type AuthorizeHandlerMiddleware struct {
}

func NewAuthorizeHandlerMiddleware() *AuthorizeHandlerMiddleware {
	return &AuthorizeHandlerMiddleware{}
}

func (m *AuthorizeHandlerMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sid := ctxdata.GetStudentIdFromCtx(r.Context())
		gradeRole := ctxdata.GetRoleGradeFromCtx(r.Context())
		licence := ctxdata.GetLicence(r.Context())
		fmt.Println(sid)
		fmt.Println(gradeRole)
		fmt.Println(licence)

	}
}
