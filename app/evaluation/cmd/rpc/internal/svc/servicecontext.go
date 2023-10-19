package svc

import (
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/config"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	InfoModel model.EvaluationInfoModel
	LikeModel model.EvaluationLikeModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewMysql(c.Mysql.Dsn())
	return &ServiceContext{
		Config:    c,
		InfoModel: model.NewEvaluationInfoModel(db),
		LikeModel: model.NewEvaluationLikeModel(db),
	}
}
