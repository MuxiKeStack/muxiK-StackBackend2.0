package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CollectionsModel = (*customCollectionsModel)(nil)

type (
	// CollectionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCollectionsModel.
	CollectionsModel interface {
		collectionsModel
	}

	customCollectionsModel struct {
		*defaultCollectionsModel
	}
)

// NewCollectionsModel returns a model for the database table.
func NewCollectionsModel(conn sqlx.SqlConn) CollectionsModel {
	return &customCollectionsModel{
		defaultCollectionsModel: newCollectionsModel(conn),
	}
}
