package rescue

import (
	"context"
	"strconv"
	"strings"
	"time"

	"management-system/common/jwtx"
	"management-system/common/response"
	"management-system/common/utils"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"
	"management-system/service/app/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthRescueInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthRescueInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthRescueInfoLogic {
	return &AuthRescueInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthRescueInfoLogic) AuthRescueInfo(req *types.AuthReq) (resp *types.AuthResp, err error) {
	if len(strings.TrimSpace(req.Mobile)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, response.Error(100, "参数错误")
	}

	userInfo, err := l.svcCtx.UsersModel.FindOneByMobile(l.ctx, req.Mobile)
	switch err {
	case nil:
	case models.ErrNotFound:
		return nil, response.Error(100, "电话号码不存在")
	default:
		return nil, err
	}

	if userInfo.Password != utils.Md5ByString(req.Password) {
		return nil, response.Error(100, "用户密码不正确")
	}

	// ---start---
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	jwtToken, err := jwtx.GetJwtToken(l.svcCtx.Config.JwtAuthForPush.AccessSecret, now, l.svcCtx.Config.JwtAuthForPush.AccessExpire, userInfo.Id)
	if err != nil {
		return nil, response.Error(400, "Token生成失败:"+err.Error())
	}
	// ---end---

	// 设置缓存，用于后续删除，Set()不带时间，默认永久，Setex()带时间
	key := "management-system token id" + strconv.Itoa(int(userInfo.Id))
	err = l.svcCtx.RedisClient.Setex(key, jwtToken, int(l.svcCtx.Config.JwtAuthForPush.AccessExpire))
	if err != nil {
		return nil, response.Error(500, "redis error:"+err.Error())
	}

	return &types.AuthResp{
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
