package main

import (
	"flag"
	"fmt"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/config"
	eva_infoServer "github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/server/eva_info"
	eva_likeServer "github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/server/eva_like"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/evaluation.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterEvaInfoServer(grpcServer, eva_infoServer.NewEvaInfoServer(ctx))
		pb.RegisterEvaLikeServer(grpcServer, eva_likeServer.NewEvaLikeServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
