package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CurriculasModel = (*customCurriculasModel)(nil)

type (
	// CurriculasModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCurriculasModel.
	CurriculasModel interface {
		curriculasModel
	}

	customCurriculasModel struct {
		*defaultCurriculasModel
	}
)

// NewCurriculasModel returns a model for the database table.
func NewCurriculasModel(conn sqlx.SqlConn) CurriculasModel {
	return &customCurriculasModel{
		defaultCurriculasModel: newCurriculasModel(conn),
	}
}
