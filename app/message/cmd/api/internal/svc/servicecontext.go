package svc

import (
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/cmd/api/internal/config"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/cmd/api/internal/socket"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config               config.Config
	NormalMessageModel   model.NormalMessageModel
	OfficialMessageModel model.OfficialMessageModel
	Hub                  *socket.Hub
}

func NewServiceContext(c config.Config, hub *socket.Hub) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:               c,
		NormalMessageModel:   model.NewNormalMessageModel(sqlConn),
		OfficialMessageModel: model.NewOfficialMessageModel(sqlConn),
		Hub:                  hub,
	}
}
