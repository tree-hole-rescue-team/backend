package application

import (
	"context"
	"encoding/json"

	"management-system/common/response"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditFormLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditFormLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditFormLogic {
	return &EditFormLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditFormLogic) EditForm(req *types.EditReq) (resp *types.EditResp, err error) {
	logx.Infof("userId: %v", l.ctx.Value("userId"))
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()
	// 先根据userId找到报名表
	form, err := l.svcCtx.ApplicationFormsModel.FindOneByUserId(l.ctx, userId)
	if err != nil {
		return nil, response.Error(500, err.Error())
	}
	userInfo, err := l.svcCtx.UsersModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, response.Error(500, err.Error())
	}
	if form.Status == 1 || userInfo.Role > 0 {
		return nil, response.Error(500, "你的申请表已通过！")
	}
	
	form.Mobile = req.Mobile
	form.Username = req.Username
	form.Sex = req.Sex
	form.Address = req.Address
	form.Birthday = req.Birthday
	form.Email = req.Email
    
	err = l.svcCtx.ApplicationFormsModel.Update(l.ctx, form)
	if err != nil {
		return nil, response.Error(500, err.Error())
	}
	return
}
