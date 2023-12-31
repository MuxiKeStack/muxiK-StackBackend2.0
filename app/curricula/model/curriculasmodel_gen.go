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
	curriculasFieldNames          = builder.RawFieldNames(&Curriculas{})
	curriculasRows                = strings.Join(curriculasFieldNames, ",")
	curriculasRowsExpectAutoSet   = strings.Join(stringx.Remove(curriculasFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	curriculasRowsWithPlaceHolder = strings.Join(stringx.Remove(curriculasFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	curriculasModel interface {
		Insert(ctx context.Context, data *Curriculas) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Curriculas, error)
		Update(ctx context.Context, data *Curriculas) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCurriculasModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Curriculas struct {
		Id              int64           `db:"id"`
		Cid             int64           `db:"cid"` // 课程编号
		CurriculaName   string          `db:"curricula_name"`
		Teacher         string          `db:"teacher"`
		Type            int64           `db:"type"`
		Rate            sql.NullFloat64 `db:"rate"`
		StartsNum       sql.NullInt64   `db:"starts_num"`
		GradeSampleSize sql.NullInt64   `db:"grade_sample_size"`
		TotalGrade      sql.NullFloat64 `db:"total_grade"`
		UsualGrade      sql.NullFloat64 `db:"usual_grade"`
		GradeR1         sql.NullInt64   `db:"grade_r1"`
		GradeR2         sql.NullInt64   `db:"grade_r2"`
		GradeR3         sql.NullInt64   `db:"grade_r3"`
	}
)

func newCurriculasModel(conn sqlx.SqlConn) *defaultCurriculasModel {
	return &defaultCurriculasModel{
		conn:  conn,
		table: "`curriculas`",
	}
}

func (m *defaultCurriculasModel) withSession(session sqlx.Session) *defaultCurriculasModel {
	return &defaultCurriculasModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`curriculas`",
	}
}

func (m *defaultCurriculasModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCurriculasModel) FindOne(ctx context.Context, id int64) (*Curriculas, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", curriculasRows, m.table)
	var resp Curriculas
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

func (m *defaultCurriculasModel) Insert(ctx context.Context, data *Curriculas) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, curriculasRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Cid, data.CurriculaName, data.Teacher, data.Type, data.Rate, data.StartsNum, data.GradeSampleSize, data.TotalGrade, data.UsualGrade, data.GradeR1, data.GradeR2, data.GradeR3)
	return ret, err
}

func (m *defaultCurriculasModel) Update(ctx context.Context, data *Curriculas) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, curriculasRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Cid, data.CurriculaName, data.Teacher, data.Type, data.Rate, data.StartsNum, data.GradeSampleSize, data.TotalGrade, data.UsualGrade, data.GradeR1, data.GradeR2, data.GradeR3, data.Id)
	return err
}

func (m *defaultCurriculasModel) tableName() string {
	return m.table
}
