package config

import (
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql   gormc.Mysql
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}
