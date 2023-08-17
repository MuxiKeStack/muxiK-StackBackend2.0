package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ EvaluationInfoModel = (*customEvaluationInfoModel)(nil)

type (
	// EvaluationInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEvaluationInfoModel.
	EvaluationInfoModel interface {
		evaluationInfoModel
	}

	customEvaluationInfoModel struct {
		*defaultEvaluationInfoModel
	}
)

// NewEvaluationInfoModel returns a model for the database table.
func NewEvaluationInfoModel(conn sqlx.SqlConn) EvaluationInfoModel {
	return &customEvaluationInfoModel{
		defaultEvaluationInfoModel: newEvaluationInfoModel(conn),
	}
}
