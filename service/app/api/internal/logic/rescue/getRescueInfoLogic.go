package rescue

import (
	"context"

	"management-system/common/response"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRescueInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRescueInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRescueInfoLogic {
	return &GetRescueInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRescueInfoLogic) GetRescueInfo(req *types.GetRescueInfoReq) (resp *types.GetRescueInfoResp, err error) {
	logx.Infof("userId: %v", l.ctx.Value("userId"))

	RescueInfos, err := l.svcCtx.RescueInfoModel.FindAllByPage(l.ctx, l.svcCtx.RescueInfoModel.RowBuilder(),req.Page, req.PageSize)
	if err != nil {
		return nil, response.Error(500, err.Error())
	}

	var result []types.RescueInfo
	if len(RescueInfos) > 0 {
		for _, rescueinfomodel := range RescueInfos {
			var typeRescueInfo types.RescueInfo
			_ = copier.Copy(&typeRescueInfo, rescueinfomodel)
			result = append(result, typeRescueInfo)
		}
	}
	return &types.GetRescueInfoResp{
		List: result,
	}, nil
}
