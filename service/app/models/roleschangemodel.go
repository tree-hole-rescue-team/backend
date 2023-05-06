package models

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RolesChangeModel = (*customRolesChangeModel)(nil)

type (
	// RolesChangeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRolesChangeModel.
	RolesChangeModel interface {
		rolesChangeModel
		FindAllByUserId(ctx context.Context, userId int64) ([]RolesChange, error)
		TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
	}

	customRolesChangeModel struct {
		*defaultRolesChangeModel
	}
)

// NewRolesChangeModel returns a model for the database table.
func NewRolesChangeModel(conn sqlx.SqlConn) RolesChangeModel {
	return &customRolesChangeModel{
		defaultRolesChangeModel: newRolesChangeModel(conn),
	}
}

func (m *defaultRolesChangeModel) FindAllByUserId(ctx context.Context, userId int64) ([]RolesChange, error) {
	query := fmt.Sprintf("select %s from %s where user_id = ?", rolesChangeRows, m.table) // %s为string类型或[]byte的占位符
	var resp []RolesChange
	err := m.conn.QueryRowsCtx(context.Background(), &resp, query, userId)
	// var resp []RolesChange
	// rows := "roles_change.`id`, roles_change.`create_time`,roles_change.`update_time`,roles_change.`delete_time`, roles_change.`user_id`, roles_change.`username`,roles_change.`operator_id`,roles_change.`operator_name`,roles_change.`new_role`,roles_change.`new_rolename`,roles_change.`old_role`,roles_change.`old_rolename`"
	// query := fmt.Sprintf("select %s from %s where user_id = %d", rows, m.table, userId)
	// fmt.Println(query)
	// err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, sqlx.ErrNotFound
	default:
		return nil, err
	}
}

// 事务
func (m *defaultRolesChangeModel) TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.conn.TransactCtx(ctx, func(ctx context.Context, s sqlx.Session) error {
		return fn(ctx, s)
	})
}
