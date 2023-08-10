package svc

import (
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/internal/config"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/model"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/go-redis/redis/v8"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Client
	UserModel   model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: "",
		DB:       0,
	})
	db, err := gormc.ConnectMysql(c.Mysql)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:      c,
		RedisClient: rdb,
		UserModel:   model.NewUserModel(db, c.Cache),
	}
}
