package ctxdata

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	CtxKeyJwtUserID    = "jwtStudentID"
	CtxKeyJwtRoleGrade = "jwtRoleGrade"
	CtxKeyJwtLicence   = "jwtLicence"
)

func GetStudentIdFromCtx(ctx context.Context) string {
	StudentID, ok := ctx.Value(CtxKeyJwtUserID).(string)
	if !ok {
		logx.WithContext(ctx).Errorf("GetStudentIDFromCtx failed")
	}
	return StudentID
}

func GetRoleGradeFromCtx(ctx context.Context) int64 {
	RoleGrade, ok := ctx.Value(CtxKeyJwtRoleGrade).(int64)
	if !ok {
		logx.WithContext(ctx).Errorf("GetRoleGradeFromCtx failed")
	}
	return RoleGrade
}

func GetLicence(ctx context.Context) int64 {
	Licence, ok := ctx.Value(CtxKeyJwtLicence).(int64)
	if !ok {
		logx.WithContext(ctx).Errorf("GetLicenceFromCtx failed")
	}
	return Licence
}
