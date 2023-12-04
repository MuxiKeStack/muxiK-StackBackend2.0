package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OfficialMessageModel = (*customOfficialMessageModel)(nil)

type (
	// OfficialMessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOfficialMessageModel.
	OfficialMessageModel interface {
		officialMessageModel
		FindBySid(ctx context.Context, sid string) ([]*OfficialMessage, error)
		FindByNullSid(ctx context.Context) ([]*OfficialMessage, error)
	}

	customOfficialMessageModel struct {
		*defaultOfficialMessageModel
	}
)

// NewOfficialMessageModel returns a model for the database table.
func NewOfficialMessageModel(conn sqlx.SqlConn) OfficialMessageModel {
	return &customOfficialMessageModel{
		defaultOfficialMessageModel: newOfficialMessageModel(conn),
	}
}

func (m *defaultOfficialMessageModel) FindBySid(ctx context.Context, sid string) ([]*OfficialMessage, error) {
	query := fmt.Sprintf("select %s from %s where `object_sid` = ? order by `send_at`", officialMessageRows, m.table)
	var resp []*OfficialMessage
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

func (m *defaultOfficialMessageModel) FindByNullSid(ctx context.Context) ([]*OfficialMessage, error) {
	query := fmt.Sprintf("select %s from %s where `object_sid` is null ", officialMessageRows, m.table)
	var resp []*OfficialMessage
	err := m.conn.QueryRowCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
