package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CurriculasModel = (*customCurriculasModel)(nil)

type Cinfo struct {
	CurriculaId   uint32 `json:"curriculaId"`   //课程号
	CurriculaName string `json:"curriculaName"` //课程名称
	Teacher       string `json:"teacher"`       //授课教师
	Type          uint8  `json:"type"`          //课程类型，公共课为0，专业课为1
}

type (
	// CurriculasModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCurriculasModel.
	CurriculasModel interface {
		curriculasModel
		customCurriculasLogicModel
	}

	customCurriculasModel struct {
		*defaultCurriculasModel
	}

	customCurriculasLogicModel interface {
		FindByInfos(ctx context.Context, infos Cinfo) (*Curriculas, error)
		FindMany(ctx context.Context, infos Cinfo) ([]*Curriculas, error)
		FindByTag(ctx context.Context, tag uint8) ([]*Curriculas, error)
	}
)

func (c customCurriculasModel) FindByTag(ctx context.Context, tag uint8) ([]*Curriculas, error) {
	var resp []*Curriculas
	query := fmt.Sprintf("select %s from %s where `tag` = ?", curriculasRows, c.table)
	err := c.conn.QueryRowCtx(ctx, &resp, query, tag)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customCurriculasModel) FindMany(ctx context.Context, infos Cinfo) ([]*Curriculas, error) {
	var resp []*Curriculas
	query := fmt.Sprintf("select %s from %s where `cid` = ? and `curricula_name` = ? and `teacher` = ? and `type` = ? limit 1", curriculasRows, c.table)
	err := c.conn.QueryRowCtx(ctx, &resp, query, infos.CurriculaId, infos.CurriculaName, infos.Teacher, infos.Type)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customCurriculasModel) FindByInfos(ctx context.Context, infos Cinfo) (*Curriculas, error) {
	var resp Curriculas
	query := fmt.Sprintf("select %s from %s where `cid` = ? and `curricula_name` = ? and `teacher` = ? and `type` = ? limit 1", curriculasRows, c.table)
	err := c.conn.QueryRowCtx(ctx, &resp, query, infos.CurriculaId, infos.CurriculaName, infos.Teacher, infos.Type)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewCurriculasModel returns a model for the database table.
func NewCurriculasModel(conn sqlx.SqlConn) CurriculasModel {
	return &customCurriculasModel{
		defaultCurriculasModel: newCurriculasModel(conn),
	}
}
