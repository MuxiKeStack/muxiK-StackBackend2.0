package config

import (
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql   gormc.Mysql
	Cache   cache.CacheConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Default struct {
		Avatar   string
		Username string
	}
}
