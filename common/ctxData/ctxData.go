package ctxData

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	CtxKeyJwtUserID = "jwtStudentID"
)

func GetStudentIdFromCtx(ctx context.Context) string {
	StudentID, ok := ctx.Value(CtxKeyJwtUserID).(string)
	if !ok {
		logx.WithContext(ctx).Errorf("GetStudentIDFromCtx failed")
	}
	return StudentID
}
