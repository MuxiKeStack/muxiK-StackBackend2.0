package svc

import (
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/config"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/middleware"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/curriculacenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config             config.Config
	CurriculaCenterRpc curriculacenter.Curriculacenter
	AuthorizeHandler   *middleware.AuthorizeHandlerMiddleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		CurriculaCenterRpc: curriculacenter.NewCurriculacenter(zrpc.MustNewClient(c.CurriculacenterRpcConf)),
		AuthorizeHandler:   middleware.NewAuthorizeHandlerMiddleware(),
	}
}
