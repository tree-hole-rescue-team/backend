package user

import (
	"context"
	"time"

	"management-system/common/jwtx"
	"management-system/common/response"
	"management-system/common/utils"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"
	"management-system/service/app/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// 先根据手机号查找用户是否注册
	user, err := l.svcCtx.UsersModel.FindOneByMobile(l.ctx, req.Mobile)
	if err != nil && err != models.ErrNotFound {
		return nil, response.Error(500, err.Error())
	}
	if user != nil {
		return nil, response.Error(100, "该手机号已注册")
	}

	// 创建user数据
	users := new(models.Users)
	users.Mobile = req.Mobile
	users.Password = utils.Md5ByString(req.Password)
	users.Username = req.UserName
	users.Sex = req.Sex
	users.Email = req.Email
	users.Address = req.Address
	users.Birthday = req.Birthday
	users.Rolename = "申请队员"
	if _, err := l.svcCtx.UsersModel.Insert(l.ctx, users); err != nil {
		return nil, response.Error(500, err.Error())
	}

	userInfo, err := l.svcCtx.UsersModel.FindOneByMobile(l.ctx, req.Mobile)
	if err != nil {
		return nil, response.Error(500, err.Error())
	}
	// 创建userAuth数据
	usersAuth := new(models.UsersAuth)
	usersAuth.AuthKey = req.Mobile
	usersAuth.AuthType = models.UserAuthTypeSystem
	usersAuth.UserId = userInfo.Id
	if _, err := l.svcCtx.UsersAuthModel.Insert(l.ctx, usersAuth); err != nil {
		return nil, response.Error(500, err.Error())
	}

	// 创建application_forms数据
	applicationForm := new(models.ApplicationForms)
	applicationForm.UserId = userInfo.Id
	applicationForm.Mobile = userInfo.Mobile
	applicationForm.Username = userInfo.Username
	applicationForm.Sex = userInfo.Sex
	applicationForm.Address = userInfo.Address
	applicationForm.Birthday = userInfo.Birthday
	applicationForm.Email = userInfo.Email
	if _, err := l.svcCtx.ApplicationFormsModel.Insert(l.ctx, applicationForm); err != nil {
		return nil, response.Error(500, err.Error())
	}

	userId := usersAuth.UserId
	// ---start---
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	jwtToken, err := jwtx.GetJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, l.svcCtx.Config.JwtAuth.AccessExpire, userId)
	if err != nil {
		return nil, response.Error(400, "Token生成失败:"+err.Error())
	}
	// ---end---

	return &types.RegisterResp{
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
		Role: 0,
	}, nil
}
