// Code generated by goctl. DO NOT EDIT.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	rolesChangeFieldNames          = builder.RawFieldNames(&RolesChange{})
	rolesChangeRows                = strings.Join(rolesChangeFieldNames, ",")
	rolesChangeRowsExpectAutoSet   = strings.Join(stringx.Remove(rolesChangeFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	rolesChangeRowsWithPlaceHolder = strings.Join(stringx.Remove(rolesChangeFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	rolesChangeModel interface {
		Insert(ctx context.Context, data *RolesChange) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*RolesChange, error)
		Update(ctx context.Context, data *RolesChange) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRolesChangeModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RolesChange struct {
		Id           int64     `db:"id"`
		CreateTime   time.Time `db:"create_time"`
		UpdateTime   time.Time `db:"update_time"`
		DeleteTime   time.Time `db:"delete_time"`
		UserId       int64     `db:"user_id"`       // 身份变动用户id
		Username     string    `db:"username"`      // 身份变动用户姓名
		OperatorId   int64     `db:"operator_id"`   // 操作人id
		OperatorName string    `db:"operator_name"` // 操作人姓名
		NewRole      int64     `db:"new_role"`      // 新身份
		NewRolename  string    `db:"new_rolename"`  // 新身份名称
		OldRole      int64     `db:"old_role"`      // 旧身份
		OldRolename  string    `db:"old_rolename"`  // 旧身份名称
	}
)

func newRolesChangeModel(conn sqlx.SqlConn) *defaultRolesChangeModel {
	return &defaultRolesChangeModel{
		conn:  conn,
		table: "`roles_change`",
	}
}

func (m *defaultRolesChangeModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultRolesChangeModel) FindOne(ctx context.Context, id int64) (*RolesChange, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", rolesChangeRows, m.table)
	var resp RolesChange
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

func (m *defaultRolesChangeModel) Insert(ctx context.Context, data *RolesChange) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, rolesChangeRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.DeleteTime, data.UserId, data.Username, data.OperatorId, data.OperatorName, data.NewRole, data.NewRolename, data.OldRole, data.OldRolename)
	return ret, err
}

func (m *defaultRolesChangeModel) Update(ctx context.Context, data *RolesChange) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, rolesChangeRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.DeleteTime, data.UserId, data.Username, data.OperatorId, data.OperatorName, data.NewRole, data.NewRolename, data.OldRole, data.OldRolename, data.Id)
	return err
}

func (m *defaultRolesChangeModel) tableName() string {
	return m.table
}