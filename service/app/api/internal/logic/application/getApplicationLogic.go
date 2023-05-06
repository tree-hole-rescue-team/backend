package application

import (
	"context"

	"management-system/common/response"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApplicationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApplicationLogic {
	return &GetApplicationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApplicationLogic) GetApplication(req *types.ApplicationFormsReq) (resp *types.ApplicationFormsResp, err error) {
	logx.Infof("userId: %v", l.ctx.Value("userId"))

	GetAllResp, err := l.svcCtx.ApplicationFormsModel.FindAllByPage(l.ctx, l.svcCtx.ApplicationFormsModel.RowBuilder(), req.Page, req.PageSize)
	if err != nil {
		return nil, response.Error(500, err.Error())
	}

	var result []types.ApplicationForm
	if len(GetAllResp) > 0 {
		for _, applicationformsmodel := range GetAllResp {
			var typeApplicationForm types.ApplicationForm
			typeApplicationForm.Id = applicationformsmodel.Id
			typeApplicationForm.CreateTime = applicationformsmodel.CreateTime.Format("2006-01-02 15:04:05")
			typeApplicationForm.UpdateTime = applicationformsmodel.UpdateTime.Format("2006-01-02 15:04:05")
			typeApplicationForm.UserId = applicationformsmodel.UserId
			typeApplicationForm.Mobile = applicationformsmodel.Mobile
			typeApplicationForm.Username = applicationformsmodel.Username
			typeApplicationForm.Sex = applicationformsmodel.Sex
			typeApplicationForm.Address = applicationformsmodel.Address
			typeApplicationForm.Birthday = applicationformsmodel.Birthday
			typeApplicationForm.Email = applicationformsmodel.Email
			typeApplicationForm.Status = applicationformsmodel.Status
			typeApplicationForm.Operator_id = applicationformsmodel.OperatorId
			typeApplicationForm.Operator_name = applicationformsmodel.OperatorName
			result = append(result, typeApplicationForm)
		}
	}
	return &types.ApplicationFormsResp{
		List: result,
	}, nil
}
