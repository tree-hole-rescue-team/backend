package models

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RescueInfoModel = (*customRescueInfoModel)(nil)

type (
	// RescueInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRescueInfoModel.
	RescueInfoModel interface {
		rescueInfoModel
		RowBuilder() squirrel.SelectBuilder
		FindAllByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64) ([]RescueInfo, error)
		
	}

	customRescueInfoModel struct {
		*defaultRescueInfoModel
	}
)

// NewRescueInfoModel returns a model for the database table.
func NewRescueInfoModel(conn sqlx.SqlConn) RescueInfoModel {
	return &customRescueInfoModel{
		defaultRescueInfoModel: newRescueInfoModel(conn),
	}
}

func (m *defaultRescueInfoModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(rescueInfoRows).From(m.table)
}

func (m *defaultRescueInfoModel) FindAllByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64) ([]RescueInfo, error) {
	var resp []RescueInfo
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	rows := "rescue_info.`id`, rescue_info.`weibo_account`, rescue_info.`area`, rescue_info.`sex`, rescue_info.`age`, rescue_info.`release_time`, rescue_info.`main_demand`,rescue_info.`risk_level`,rescue_info.`status`,rescue_info.`rescue_teacher_id`,rescue_info.`rescue_teacher_name`"
	query := fmt.Sprintf("select %s from %s  limit %d,%d", rows, m.table, offset, pageSize)

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
