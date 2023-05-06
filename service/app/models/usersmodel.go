package models

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		RowBuilder() squirrel.SelectBuilder
		// FindAll(ctx context.Context, where string, Page, PageSize int64) ([]Users, error)
		// FindAllByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Users, error)
		FindAllByPage(ctx context.Context, page, pageSize int64, orderBy string) ([]Users, error)
		FindUserByRole(ctx context.Context, role int64) ([]Users, error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn),
	}
}

// export logic
func (m *defaultUsersModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(usersRows).From(m.table)
}

// 查看全部role=0的用户信息
func (m *defaultUsersModel) FindUserByRole(ctx context.Context, role int64) ([]Users, error) {
	var resp []Users
	rows := "users.`id`, users.`create_time`, users.`update_time`, users.`delete_time`, users.`open_id`, users.`mobile`, users.`username`,users.`password`,users.`sex`,users.`email`,users.`role`,users.`rolename`,users.`avatar`,users.`address`,users.`birthday`"
	query := fmt.Sprintf("select %s from %s where role = %d", rows, m.table, role)
	fmt.Println(query)
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return resp, ErrNotFound
	default:
		return nil, err
	}
}

// 根据页数和每页数量查找role=0的用户信息
func (m *defaultUsersModel) FindAllByPage(ctx context.Context, page, pageSize int64, where string) ([]Users, error) {
	var resp []Users
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	rows := "users.`id`, users.`create_time`, users.`update_time`, users.`delete_time`, users.`open_id`, users.`mobile`, users.`username`,users.`password`,users.`sex`,users.`email`,users.`role`,users.`rolename`,users.`avatar`,users.`address`,users.`birthday`"
	query := fmt.Sprintf("select %s from %s where %s  order by id  limit %d,%d", rows, m.table, where, offset, pageSize)
	fmt.Println(query)
	err := m.conn.QueryRowsCtx(ctx, &resp, query) // 这里报错"unsupported unmarshal type"，注意是QueryRowsCtx，不是QueryRowCtx
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
// csdn版本
/* func (m *defaultUsersModel) FindAll(ctx context.Context, where string, Page, PageSize int64) ([]Users, error) {
	var resp []Users
	var err error
	offset := (Page - 1) * PageSize
	if len(where) == 0 {
		query := fmt.Sprintf("select %s from %s order by create_time desc limit ?,?", "*", m.table)
		err = m.conn.QueryRowCtx(ctx, &resp, query, offset, PageSize)
	} else {
		query := fmt.Sprintf("select %s from %s where %s order by create_time desc limit ?,?", "*", m.table, where)
		err = m.conn.QueryRowCtx(ctx, &resp, query, offset, PageSize)
	}
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
} */

// looklook版本 报错"unsupported unmarshal type"
/* func (m *defaultUsersModel) FindAllByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, where string) ([]*Users, error) {

	rowBuilder = rowBuilder.OrderBy("id DESC")

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := rowBuilder.Where(where).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Users
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
} */

