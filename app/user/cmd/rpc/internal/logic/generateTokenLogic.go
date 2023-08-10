package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ctxdata"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/xerr"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"time"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.GetStudentID(), in.GetRoleGrade(), in.GetLicence())
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "getJwtToken err userId:%d , err:%v", in.StudentID, err)
	}

	return &pb.GenerateTokenResponse{
		Token: accessToken,
	}, nil
}

func (l *GenerateTokenLogic) getJwtToken(secretKey string, iat, seconds int64, sid string, roleGrade int64, licence int64) (string, error) {

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat

	claims[ctxdata.CtxKeyJwtUserID] = sid
	claims[ctxdata.CtxKeyJwtRoleGrade] = roleGrade
	claims[ctxdata.CtxKeyJwtLicence] = licence
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
