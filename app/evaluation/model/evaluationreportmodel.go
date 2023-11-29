package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ EvaluationReportModel = (*customEvaluationReportModel)(nil)

type (
	// EvaluationReportModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEvaluationReportModel.
	EvaluationReportModel interface {
		evaluationReportModel
		ReportLogic
	}

	customEvaluationReportModel struct {
		*defaultEvaluationReportModel
	}
	ReportLogic interface {
		FindReport(ctx context.Context, id string, method int) (resp *EvaluationReport, err error)
	}
)

// NewEvaluationReportModel returns a model for the database table.
func NewEvaluationReportModel(conn sqlx.SqlConn) EvaluationReportModel {
	return &customEvaluationReportModel{
		defaultEvaluationReportModel: newEvaluationReportModel(conn),
	}
}

func (m *customEvaluationReportModel) FindReport(ctx context.Context, id string, method int) (resp *EvaluationReport, err error) {
	name := ""
	switch method {
	case 1:
		name = "pid"
	case 2:
		name = "sid"
	case 3:
		name = "cid"
	case 4:
		name = "rid"
	}
	query := fmt.Sprintf("select * from %s where `%s` = ? limit 1", m.table, name)
	err = m.conn.QueryRowCtx(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
}
