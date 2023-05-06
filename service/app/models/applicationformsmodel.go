package models

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ApplicationFormsModel = (*customApplicationFormsModel)(nil)

type (
	// ApplicationFormsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customApplicationFormsModel.
	ApplicationFormsModel interface {
		applicationFormsModel
		RowBuilder() squirrel.SelectBuilder
		TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
		FindAllByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64) ([]ApplicationForms, error)
		FindByStatus(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, where string) ([]ApplicationForms, error)
	}

	customApplicationFormsModel struct {
		*defaultApplicationFormsModel
	}
)

// NewApplicationFormsModel returns a model for the database table.
func NewApplicationFormsModel(conn sqlx.SqlConn) ApplicationFormsModel {
	return &customApplicationFormsModel{
		defaultApplicationFormsModel: newApplicationFormsModel(conn),
	}
}

func (m *defaultApplicationFormsModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(applicationFormsRows).From(m.table)
}

// 事务
func (m *defaultApplicationFormsModel) TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.conn.TransactCtx(ctx, func(ctx context.Context, s sqlx.Session) error {
		return fn(ctx, s)
	})
}

func (m *defaultApplicationFormsModel) FindAllByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64) ([]ApplicationForms, error) {

	rowBuilder = rowBuilder.OrderBy("update_time DESC") // 按更新时间排序，先查询到最先更新的

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	query, values, err := rowBuilder.Offset(uint64(offset)).Limit(uint64(pageSize)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []ApplicationForms
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultApplicationFormsModel) FindByStatus(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, where string) ([]ApplicationForms, error) {

	rowBuilder = rowBuilder.OrderBy("update_time DESC") // 按更新时间排序，先查询到最先更新的

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	query, values, err := rowBuilder.Where(where).Offset(uint64(offset)).Limit(uint64(pageSize)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []ApplicationForms
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
