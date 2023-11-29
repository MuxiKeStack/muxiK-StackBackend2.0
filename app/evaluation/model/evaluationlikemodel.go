package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ EvaluationLikeModel = (*customEvaluationLikeModel)(nil)

type (
	// EvaluationLikeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEvaluationLikeModel.
	EvaluationLikeModel interface {
		evaluationLikeModel
		LikeLogic
	}

	customEvaluationLikeModel struct {
		*defaultEvaluationLikeModel
	}

	LikeLogic interface {
		FindLike(ctx context.Context, pid string, sid string) (resp *EvaluationLike, err error)
	}
)

// NewEvaluationLikeModel returns a model for the database table.
func NewEvaluationLikeModel(conn sqlx.SqlConn) EvaluationLikeModel {
	return &customEvaluationLikeModel{
		defaultEvaluationLikeModel: newEvaluationLikeModel(conn),
	}
}

func (m *customEvaluationLikeModel) FindLike(ctx context.Context, pid string, sid string) (resp *EvaluationLike, err error) {
	query := fmt.Sprintf("select * from %s where `pid` = ? and `sid` = ? limit 1", m.table)
	err = m.conn.QueryRowCtx(ctx, &resp, query, pid, sid)
	if err != nil {
		return nil, err
	}
	return
}
