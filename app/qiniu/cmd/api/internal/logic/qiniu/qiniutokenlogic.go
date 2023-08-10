package qiniu

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/qiniu/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/qiniu/cmd/api/internal/types"
	qiniuToken "github.com/MuxiKeStack/muxiK-StackBackend2.0/common/qiniu"
	"github.com/zeromicro/go-zero/core/logx"
)

type Qiniu_tokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQiniu_tokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Qiniu_tokenLogic {
	return &Qiniu_tokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Qiniu_tokenLogic) Qiniu_token(req *types.QiniuTokenRequest) (resp *types.QiniuTokenResponse, err error) {
	token := qiniuToken.GetToken(l.svcCtx.Config.OSS.AccessKey,
		l.svcCtx.Config.OSS.SecretKey,
		l.svcCtx.Config.OSS.BucketName,
		l.svcCtx.Config.OSS.DomainName)

	return &types.QiniuTokenResponse{
		Token: token,
	}, nil
}
