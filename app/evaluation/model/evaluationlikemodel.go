package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ EvaluationLikeModel = (*customEvaluationLikeModel)(nil)

type (
	// EvaluationLikeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEvaluationLikeModel.
	EvaluationLikeModel interface {
		evaluationLikeModel
	}

	customEvaluationLikeModel struct {
		*defaultEvaluationLikeModel
	}
)

// NewEvaluationLikeModel returns a model for the database table.
func NewEvaluationLikeModel(conn sqlx.SqlConn) EvaluationLikeModel {
	return &customEvaluationLikeModel{
		defaultEvaluationLikeModel: newEvaluationLikeModel(conn),
	}
}
