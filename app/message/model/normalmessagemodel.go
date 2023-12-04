package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NormalMessageModel = (*customNormalMessageModel)(nil)

type (
	// NormalMessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNormalMessageModel.
	NormalMessageModel interface {
		normalMessageModel
		FindBySid(ctx context.Context, sid string) ([]*NormalMessage, error)
	}

	customNormalMessageModel struct {
		*defaultNormalMessageModel
	}
)

// NewNormalMessageModel returns a model for the database table.
func NewNormalMessageModel(conn sqlx.SqlConn) NormalMessageModel {
	return &customNormalMessageModel{
		defaultNormalMessageModel: newNormalMessageModel(conn),
	}
}

func (m *customNormalMessageModel) FindBySid(ctx context.Context, sid string) ([]*NormalMessage, error) {
	query := fmt.Sprintf("select %s from %s where `object_sid` = ? order by `send_at`", normalMessageRows, m.table)
	var resp []*NormalMessage
	err := m.conn.QueryRowsCtx(ctx, &resp, query, sid)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
