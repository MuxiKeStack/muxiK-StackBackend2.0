package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CollectionsModel = (*customCollectionsModel)(nil)

type CollectionInfo struct {
	Sid string
	Cid int64
}

type (
	// CollectionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCollectionsModel.
	CollectionsModel interface {
		collectionsModel
		collectionLogicModel
	}

	customCollectionsModel struct {
		*defaultCollectionsModel
	}

	collectionLogicModel interface {
		FindByInfos(ctx context.Context, infos CollectionInfo) (*Collections, error)
	}
)

func (c customCollectionsModel) FindByInfos(ctx context.Context, infos CollectionInfo) (*Collections, error) {
	//TODO implement me
	var resp Collections
	query := fmt.Sprintf("select %s from %s where `sid` = ? and `c_data_id` = ? limit 1", curriculasRows, c.table)
	err := c.conn.QueryRowCtx(ctx, &resp, query, infos.Sid, infos.Cid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewCollectionsModel returns a model for the database table.
func NewCollectionsModel(conn sqlx.SqlConn) CollectionsModel {
	return &customCollectionsModel{
		defaultCollectionsModel: newCollectionsModel(conn),
	}
}
