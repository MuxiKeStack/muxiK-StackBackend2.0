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
		IncreaseIntegral(ctx context.Context, sid string, integral int64) (*User, error)
	}
)

func (c customUserModel) IncreaseIntegral(ctx context.Context, sid string, integral int64) (*User, error) {
	var user *User
	err := c.QueryNoCacheCtx(ctx, &user, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&User{}).Where("`sid` = ?", sid).First(&user).Error
	})
	if err != nil {
		return nil, err
	}
	user.Integral += integral
	switch user.Integral {
	case 5:
		user.RoleGrade = 1
	case 20:
		user.RoleGrade = 2
	case 80:
		user.RoleGrade = 3
	case 150:
		user.RoleGrade = 4
	case 300:
		user.RoleGrade = 5
	}
	err = c.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		return db.Save(&user).Error
	}, c.getCacheKeys(user)...)
	return user, err
}

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
