package application

import (
	"context"
	"encoding/json"

	"management-system/common/response"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyApplicationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyApplicationLogic {
	return &GetMyApplicationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyApplicationLogic) GetMyApplication(req *types.ApplicationFormReq) (resp *types.ApplicationFormResp, err error) {
	logx.Infof("userId: %v", l.ctx.Value("userId"))
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()

	form, err := l.svcCtx.ApplicationFormsModel.FindOneByUserId(l.ctx, userId)
	if err != nil {
		return nil, response.Error(500, err.Error())
	}

	var typeApplicationForm types.ApplicationForm
	typeApplicationForm.Id = form.Id
	typeApplicationForm.CreateTime = form.CreateTime.Format("2006-01-02 15:04:05")
	typeApplicationForm.UpdateTime = form.UpdateTime.Format("2006-01-02 15:04:05")
	typeApplicationForm.UserId = form.UserId
	typeApplicationForm.Mobile = form.Mobile
	typeApplicationForm.Username = form.Username
	typeApplicationForm.Sex = form.Sex
	typeApplicationForm.Address = form.Address
	typeApplicationForm.Birthday = form.Birthday
	typeApplicationForm.Email = form.Email
	typeApplicationForm.Status = form.Status
	typeApplicationForm.Operator_id = form.OperatorId
	typeApplicationForm.Operator_name = form.OperatorName

	return &types.ApplicationFormResp{
		Form: typeApplicationForm,
	}, nil
}
