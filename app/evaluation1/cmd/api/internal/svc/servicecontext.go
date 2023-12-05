package svc

import (
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation1/cmd/api/internal/config"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation1/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	CommentModel model.CommentModel
	ReplyModel   model.ReplyModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:       c,
		CommentModel: model.NewCommentModel(sqlConn),
		ReplyModel:   model.NewReplyModel(sqlConn),
	}
}
