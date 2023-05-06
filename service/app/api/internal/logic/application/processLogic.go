package application

import (
	"context"
	"encoding/json"
	"time"

	"management-system/common/response"
	"management-system/common/role"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"
	"management-system/service/app/models"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ProcessLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcessLogic {
	return &ProcessLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcessLogic) Process(req *types.ProcessReq) (resp *types.ProcessResp, err error) {
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()
	// 操作人信息
	operatorInfo, err := l.svcCtx.UsersModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, response.Error(500, err.Error())
	}
	if operatorInfo.Role < 32 {
		return nil, response.Error(403, "权限不够")
	}
	// 申请表信息
	applicationInfo, err := l.svcCtx.ApplicationFormsModel.FindOne(l.ctx, req.ApplicationFormId)
	if err != nil {
		return nil, response.Error(500, err.Error())
	}
	// 修改申请表信息
	applicationInfo.Status = req.Status
	applicationInfo.OperatorId = userId
	applicationInfo.OperatorName = operatorInfo.Username
	// 申请人信息
	applicantInfo, err := l.svcCtx.UsersModel.FindOne(l.ctx, applicationInfo.UserId)
	// 创建roles_change表数据
	rolesChange := new(models.RolesChange)
	rolesChange.UserId = applicantInfo.Id
	rolesChange.Username = applicantInfo.Username
	rolesChange.OperatorId = userId
	rolesChange.OperatorName = operatorInfo.Username
	rolesChange.NewRole = 1
	rolesChange.NewRolename = role.GetRoleName(1)
	rolesChange.OldRole = 0
	rolesChange.OldRolename = role.GetRoleName(0)
	// 修改userInfo数据
	applicantInfo.Role = 1
	applicantInfo.Rolename = role.GetRoleName(1)

	if err := l.svcCtx.ApplicationFormsModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		if _, err := l.svcCtx.RolesChangeModel.Insert(l.ctx, rolesChange); err != nil {
			return err
		}
		if err := l.svcCtx.UsersModel.Update(l.ctx, applicantInfo); err != nil {
			return err
		}
		if err := l.svcCtx.ApplicationFormsModel.Update(l.ctx, applicationInfo); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, response.Error(500, err.Error())
	}

	return &types.ProcessResp{
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}