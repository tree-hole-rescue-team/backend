package application

import (
	"context"
	"fmt"

	"management-system/common/response"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByStatusLogic {
	return &GetByStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByStatusLogic) GetByStatus(req *types.GetByStatusReq) (resp *types.GetByStatusResp, err error) {
	logx.Infof("userId: %v", l.ctx.Value("userId"))

	where := fmt.Sprintf("status = %d", req.Status)
	GetAllResp, err := l.svcCtx.ApplicationFormsModel.FindByStatus(l.ctx, l.svcCtx.ApplicationFormsModel.RowBuilder(), req.Page, req.PageSize, where)
	if err != nil {
		return nil, response.Error(500, err.Error())
	}

	var result []types.ApplicationForm
	if len(GetAllResp) > 0 {
		for _, applicationformsmodel := range GetAllResp {
			var typeUser types.ApplicationForm
			typeUser.Id = applicationformsmodel.Id
			typeUser.CreateTime = applicationformsmodel.CreateTime.Format("2006-01-02 15:04:05")
			typeUser.UpdateTime = applicationformsmodel.UpdateTime.Format("2006-01-02 15:04:05")
			typeUser.UserId = applicationformsmodel.UserId
			typeUser.Mobile = applicationformsmodel.Mobile
			typeUser.Username = applicationformsmodel.Username
			typeUser.Sex = applicationformsmodel.Sex
			typeUser.Address = applicationformsmodel.Address
			typeUser.Birthday = applicationformsmodel.Birthday
			typeUser.Email = applicationformsmodel.Email
			typeUser.Status = applicationformsmodel.Status
			typeUser.Operator_id = applicationformsmodel.OperatorId
			typeUser.Operator_name = applicationformsmodel.OperatorName
			result = append(result, typeUser)
		}
	}
	return &types.GetByStatusResp{
		List: result,
	}, nil
}
