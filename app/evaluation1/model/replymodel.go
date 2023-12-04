package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ReplyModel = (*customReplyModel)(nil)

type (
	// ReplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customReplyModel.
	ReplyModel interface {
		replyModel
	}

	customReplyModel struct {
		*defaultReplyModel
	}
)

// NewReplyModel returns a model for the database table.
func NewReplyModel(conn sqlx.SqlConn) ReplyModel {
	return &customReplyModel{
		defaultReplyModel: newReplyModel(conn),
	}
}
