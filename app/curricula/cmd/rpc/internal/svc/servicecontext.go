package svc

import (
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/config"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	CurriculaModel model.CurriculasModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		CurriculaModel: model.NewCurriculasModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
