package main

import (
	"flag"
	"fmt"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/rpc/internal/config"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/rpc/internal/server"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/userCenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUsercenterServer(grpcServer, server.NewUsercenterServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
