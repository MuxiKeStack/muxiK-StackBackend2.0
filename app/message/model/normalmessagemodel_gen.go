// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	normalMessageFieldNames          = builder.RawFieldNames(&NormalMessage{})
	normalMessageRows                = strings.Join(normalMessageFieldNames, ",")
	normalMessageRowsExpectAutoSet   = strings.Join(stringx.Remove(normalMessageFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	normalMessageRowsWithPlaceHolder = strings.Join(stringx.Remove(normalMessageFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	normalMessageModel interface {
		Insert(ctx context.Context, data *NormalMessage) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*NormalMessage, error)
		Update(ctx context.Context, data *NormalMessage) error
		Delete(ctx context.Context, id int64) error
	}

	defaultNormalMessageModel struct {
		conn  sqlx.SqlConn
		table string
	}

	NormalMessage struct {
		Id              int64          `db:"id"`
		Type            string         `db:"type"`
		SenderSid       string         `db:"sender_sid"`
		ObjectSid       string         `db:"object_sid"`
		OriginContentId int64          `db:"origin_content_id"`
		Text            sql.NullString `db:"text"`
		SendAt          int64          `db:"send_at"`
	}
)

func newNormalMessageModel(conn sqlx.SqlConn) *defaultNormalMessageModel {
	return &defaultNormalMessageModel{
		conn:  conn,
		table: "`normal_message`",
	}
}

func (m *defaultNormalMessageModel) withSession(session sqlx.Session) *defaultNormalMessageModel {
	return &defaultNormalMessageModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`normal_message`",
	}
}

func (m *defaultNormalMessageModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultNormalMessageModel) FindOne(ctx context.Context, id int64) (*NormalMessage, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", normalMessageRows, m.table)
	var resp NormalMessage
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultNormalMessageModel) Insert(ctx context.Context, data *NormalMessage) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, normalMessageRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Type, data.SenderSid, data.ObjectSid, data.OriginContentId, data.Text, data.SendAt)
	return ret, err
}

func (m *defaultNormalMessageModel) Update(ctx context.Context, data *NormalMessage) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, normalMessageRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Type, data.SenderSid, data.ObjectSid, data.OriginContentId, data.Text, data.SendAt, data.Id)
	return err
}

func (m *defaultNormalMessageModel) tableName() string {
	return m.table
}