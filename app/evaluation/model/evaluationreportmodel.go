package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ EvaluationReportModel = (*customEvaluationReportModel)(nil)

type (
	// EvaluationReportModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEvaluationReportModel.
	EvaluationReportModel interface {
		evaluationReportModel
	}

	customEvaluationReportModel struct {
		*defaultEvaluationReportModel
	}
)

// NewEvaluationReportModel returns a model for the database table.
func NewEvaluationReportModel(conn sqlx.SqlConn) EvaluationReportModel {
	return &customEvaluationReportModel{
		defaultEvaluationReportModel: newEvaluationReportModel(conn),
	}
}
