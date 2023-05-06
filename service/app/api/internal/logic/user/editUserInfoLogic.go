package user

import (
	"context"
	"encoding/json"

	"management-system/common/response"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUserInfoLogic {
	return &EditUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditUserInfoLogic) EditUserInfo(req *types.EditUserInfoReq) (resp *types.EditUserInfoResp, err error) {
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()
	// 先根据userId找到报名表
	userInfo, err := l.svcCtx.UsersModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, response.Error(500, err.Error())
	}

	if userInfo.Role < 1 {
		return nil, response.Error(200, "您的申请表还未通过，请修改申请表！")
	}

	userInfo.Mobile = req.Mobile
	userInfo.Username = req.Username
	userInfo.Sex = req.Sex
	userInfo.Address = req.Address
	userInfo.Birthday = req.Birthday
	userInfo.Email = req.Email

	err = l.svcCtx.UsersModel.Update(l.ctx, userInfo)
	if err != nil {
		return nil,response.Error(500,err.Error())
	}
	return
}
