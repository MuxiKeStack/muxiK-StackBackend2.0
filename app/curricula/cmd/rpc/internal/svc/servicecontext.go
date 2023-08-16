package svc

import "github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
