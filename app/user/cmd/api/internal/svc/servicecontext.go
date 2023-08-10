package svc

import (
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/config"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/api/internal/middleware"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/usercenter"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	UserCenterRpc    usercenter.Usercenter
	AuthorizeHandler rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		UserCenterRpc:    usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		AuthorizeHandler: middleware.NewAuthorizeHandlerMiddleware().Handle,
	}
}
