package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		customUserLogicModel
	}

	customUserModel struct {
		*defaultUserModel
	}

	customUserLogicModel interface {
		IfExist(ctx context.Context, sid string) int64
	}
)

func (c customUserModel) IfExist(ctx context.Context, sid string) int64 {
	var num int64
	var user User
	err := c.QueryNoCacheCtx(ctx, &user, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&User{}).Where("`sid` = ?", sid).Count(&num).Error
	})
	if err != nil || num == 0 {
		return 0
	}
	return 1
}

// NewUserModel returns a model for the database table.
func NewUserModel(conn *gorm.DB, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

func (m *defaultUserModel) customCacheKeys(data *User) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
