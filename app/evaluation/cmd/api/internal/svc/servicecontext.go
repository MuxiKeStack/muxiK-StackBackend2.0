package svc

import (
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/api/internal/config"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/client/eva_info"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/client/eva_like"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/client/eva_report"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	InfoRpc   eva_info.EvaInfo
	LikeRpc   eva_like.EvaLike
	ReportRpc eva_report.EvaReport
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		InfoRpc: eva_info.NewEvaInfo(zrpc.MustNewClient(c.InfoRpcConf)),
		LikeRpc: eva_like.NewEvaLike(zrpc.MustNewClient(c.LikeRpcConf)),
	}
}
