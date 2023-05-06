package rescue

import (
	"context"

	"management-system/common/response"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"
	"management-system/service/app/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushRescueInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPushRescueInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushRescueInfoLogic {
	return &PushRescueInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PushRescueInfoLogic) PushRescueInfo(req *types.PushReq) (resp *types.PushResp, err error) {
	logx.Infof("userId: %v", l.ctx.Value("userId"))
	
	rescueInfo := new(models.RescueInfo)
	rescueInfo.WeiboAccount = req.WeiboAccount
	rescueInfo.Area = req.Area
	rescueInfo.Sex = req.Sex
	rescueInfo.Age = req.Age
	rescueInfo.ReleaseTime = req.ReleaseTime
	rescueInfo.MainDemand = req.MainDemand
	rescueInfo.RiskLevel = req.RiskLevel

	if _, err := l.svcCtx.RescueInfoModel.Insert(l.ctx, rescueInfo); err != nil {
		return nil, response.Error(500, err.Error())
	}

	return
}
