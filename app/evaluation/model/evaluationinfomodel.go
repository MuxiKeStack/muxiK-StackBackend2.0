package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ EvaluationInfoModel = (*customEvaluationInfoModel)(nil)

type (
	// EvaluationInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEvaluationInfoModel.
	EvaluationInfoModel interface {
		evaluationInfoModel
		LogicEvaluationInfoModel
	}

	customEvaluationInfoModel struct {
		*defaultEvaluationInfoModel
	}
	LogicEvaluationInfoModel interface {
		CreateEvaluationInfo(ctx context.Context, evaluationInfo *EvaluationInfo) error
		UpdateEvaluationInfo(ctx context.Context, evaluationInfo *EvaluationInfo) error
		GetEvaluationByCid(ctx context.Context, cid string) (resp []EvaluationInfo, err error)
		GetEvaluationBySid(ctx context.Context, sid string) (resp []EvaluationInfo, err error)
		ModifyLike(ctx context.Context, pid int64, delta int64) (err error)
	}
)

// NewEvaluationInfoModel returns a model for the database table.
func NewEvaluationInfoModel(conn sqlx.SqlConn) EvaluationInfoModel {
	return &customEvaluationInfoModel{
		defaultEvaluationInfoModel: newEvaluationInfoModel(conn),
	}
}

func (m *customEvaluationInfoModel) CreateEvaluationInfo(ctx context.Context, evaluationInfo *EvaluationInfo) error {
	_, err := m.Insert(ctx, evaluationInfo)
	return err
}

func (m *customEvaluationInfoModel) UpdateEvaluationInfo(ctx context.Context, evaluationInfo *EvaluationInfo) error {
	err := m.Update(ctx, evaluationInfo)
	return err
}

func (m *customEvaluationInfoModel) GetEvaluationByCid(ctx context.Context, cid string) (resp []EvaluationInfo, err error) {
	query := fmt.Sprintf("select * from %s where `sid` = ? limit 1000", m.table)
	err = m.conn.QueryRowCtx(ctx, &resp, query, cid)
	if err != nil {
		resp = nil
	}
	return
}

func (m *customEvaluationInfoModel) GetEvaluationBySid(ctx context.Context, sid string) (resp []EvaluationInfo, err error) {
	query := fmt.Sprintf("select * from %s where `sid` = ? limit 1000", m.table)
	err = m.conn.QueryRowCtx(ctx, &resp, query, sid)
	if err != nil {
		resp = nil
	}
	return
}

func (m *customEvaluationInfoModel) ModifyLike(ctx context.Context, pid int64, delta int64) (err error) {
	info, err := m.FindOne(ctx, pid)
	if err != nil {
		return err
	}
	info.Liked.Int64 = info.Liked.Int64 + delta
	err = m.Update(ctx, info)
	return
}
