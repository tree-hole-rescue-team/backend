package role

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

type ChangeRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeRoleLogic {
	return &ChangeRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeRoleLogic) ChangeRole(req *types.ChangeRoleReq) (resp *types.ChangeRoleResp, err error) {
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()
	// 被修改身份成员信息
	userInfo, err := l.svcCtx.UsersModel.FindOne(l.ctx, req.UserId)
	if err != nil {
		return nil, response.Error(500, err.Error())
	}
	// 操作人信息
	operatorInfo, err := l.svcCtx.UsersModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, response.Error(500, err.Error())
	}
	// 创建roles_change表数据
	rolesChange := new(models.RolesChange)
	rolesChange.UserId = req.UserId
	rolesChange.Username = userInfo.Username
	rolesChange.OperatorId = userId
	rolesChange.OperatorName = operatorInfo.Username
	rolesChange.NewRole = req.NewRole
	rolesChange.NewRolename = role.GetRoleName(req.NewRole)
	rolesChange.OldRole = userInfo.Role
	rolesChange.OldRolename = role.GetRoleName(rolesChange.OldRole)
	// 修改userInfo数据
	userInfo.Role = req.NewRole
	userInfo.Rolename = role.GetRoleName(req.NewRole)
	// 身份代号 0-申请队员 1-岗前培训 2-见习队员 3-正式队员 4-督导老师 30-普通队员 31-核心队员 32-区域负责人 33-组委会成员 34-组委会主任
	/* if operatorInfo.Role >= 32 {
		if _, err := l.svcCtx.RolesChangeModel.Insert(l.ctx, rolesChange); err != nil {
			return nil, response.Error(500, err.Error())
		}
		if err := l.svcCtx.UsersModel.Update(l.ctx, userInfo); err != nil {
			return nil, response.Error(500, err.Error())
		}
		return &types.ChangeRoleResp{
			CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
			NewRole:     req.NewRole,
			NewRoleName: role.GetRoleName(req.NewRole),
			OldRole:     rolesChange.OldRole,
			OldRoleName: rolesChange.OldRolename,
		}, nil
	} else {
		return nil, response.Error(403, "权限不够")
	} */
	if operatorInfo.Role >= 32 {
		if err := l.svcCtx.RolesChangeModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			if _, err := l.svcCtx.RolesChangeModel.Insert(l.ctx, rolesChange); err != nil {
				return err
			}
			if err := l.svcCtx.UsersModel.Update(l.ctx, userInfo); err != nil {
				return err
			}
			return nil
		}); err != nil {
			return nil, response.Error(500, err.Error())
		}

		return &types.ChangeRoleResp{
			CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
			NewRole:     req.NewRole,
			NewRoleName: role.GetRoleName(req.NewRole),
			OldRole:     rolesChange.OldRole,
			OldRoleName: rolesChange.OldRolename,
		}, nil
	} else {
		return nil, response.Error(403, "权限不够")
	}
}
