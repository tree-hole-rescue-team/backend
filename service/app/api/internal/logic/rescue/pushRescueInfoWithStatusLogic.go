package rescue

import (
	"context"
	"fmt"

	"management-system/common/response"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"
	"management-system/service/app/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushRescueInfoWithStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPushRescueInfoWithStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushRescueInfoWithStatusLogic {
	return &PushRescueInfoWithStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PushRescueInfoWithStatusLogic) PushRescueInfoWithStatus(req *types.PushWithStatusReq) (resp *types.PushWithStatusResp, err error) {
	logx.Infof("userId: %v", l.ctx.Value("userId"))

	rescueInfo := new(models.RescueInfo)
	rescueInfo.WeiboAccount = req.WeiboAccount
	fmt.Sprintln(req.WeiboAccount)
	rescueInfo.Area = req.Area
	rescueInfo.Sex = req.Sex
	rescueInfo.Age = req.Age
	rescueInfo.ReleaseTime = req.ReleaseTime
	rescueInfo.MainDemand = req.MainDemand
	rescueInfo.RiskLevel = req.RiskLevel
	rescueInfo.Status = int64(req.Status)

	if _, err := l.svcCtx.RescueInfoModel.Insert(l.ctx, rescueInfo); err != nil {
		return nil, response.Error(500, err.Error())
	}

	return
}
