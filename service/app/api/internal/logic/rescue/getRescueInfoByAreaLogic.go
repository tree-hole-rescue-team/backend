package rescue

import (
	"context"
	"strings"

	"management-system/common/response"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRescueInfoByAreaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRescueInfoByAreaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRescueInfoByAreaLogic {
	return &GetRescueInfoByAreaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRescueInfoByAreaLogic) GetRescueInfoByArea(req *types.GetRescueInfoByAreaReq) (resp *types.GetRescueInfoByAreaResp, err error) {
	logx.Infof("userId: %v", l.ctx.Value("userId"))

	RescueInfos, err := l.svcCtx.RescueInfoModel.FindAllByPage(l.ctx, l.svcCtx.RescueInfoModel.RowBuilder(), req.Page, req.PageSize)
	if err != nil {
		return nil, response.Error(500, err.Error())
	}

	var result []types.RescueInfo
	if len(RescueInfos) > 0 {
		for _, rescueinfomodel := range RescueInfos {
			var typeRescueInfo types.RescueInfo
			_ = copier.Copy(&typeRescueInfo, rescueinfomodel)
			if strings.Contains(typeRescueInfo.Area,req.Area){ // 匹配地区
				result = append(result, typeRescueInfo)
			}
		}
	}
	return &types.GetRescueInfoByAreaResp{
		List: result,
	}, nil
}
