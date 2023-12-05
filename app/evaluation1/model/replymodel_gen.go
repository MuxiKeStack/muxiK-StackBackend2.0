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
	replyFieldNames          = builder.RawFieldNames(&Reply{})
	replyRows                = strings.Join(replyFieldNames, ",")
	replyRowsExpectAutoSet   = strings.Join(stringx.Remove(replyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	replyRowsWithPlaceHolder = strings.Join(stringx.Remove(replyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	replyModel interface {
		Insert(ctx context.Context, data *Reply) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Reply, error)
		Update(ctx context.Context, data *Reply) error
		Delete(ctx context.Context, id int64) error
	}

	defaultReplyModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Reply struct {
		Id        int64  `db:"id"`
		CommentId int64  `db:"comment_id"` // 关联评论的comment_id
		Sid       string `db:"sid"`        // 回复者学号
		Text      string `db:"text"`       // 回复内容
		Date      int64  `db:"date"`       // 回复日期
	}
)

func newReplyModel(conn sqlx.SqlConn) *defaultReplyModel {
	return &defaultReplyModel{
		conn:  conn,
		table: "`reply`",
	}
}

func (m *defaultReplyModel) withSession(session sqlx.Session) *defaultReplyModel {
	return &defaultReplyModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`reply`",
	}
}

func (m *defaultReplyModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultReplyModel) FindOne(ctx context.Context, id int64) (*Reply, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", replyRows, m.table)
	var resp Reply
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

func (m *defaultReplyModel) Insert(ctx context.Context, data *Reply) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, replyRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.CommentId, data.Sid, data.Text, data.Date)
	return ret, err
}

func (m *defaultReplyModel) Update(ctx context.Context, data *Reply) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, replyRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.CommentId, data.Sid, data.Text, data.Date, data.Id)
	return err
}

func (m *defaultReplyModel) tableName() string {
	return m.table
}