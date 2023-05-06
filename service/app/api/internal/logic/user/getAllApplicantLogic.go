package user

import (
	"context"

	"management-system/common/response"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllApplicantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllApplicantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllApplicantLogic {
	return &GetAllApplicantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllApplicantLogic) GetAllApplicant(req *types.GetAllApplicantReq) (resp *types.GetAllApplicantResp, err error) {
	logx.Infof("userId: %v", l.ctx.Value("userId")) // 这里的key和生成jwt token时传入的key一致
	// userId, _ := l.ctx.Value("userId").(json.Number).Int64()

	// GetAllApplicantResp, err := l.svcCtx.UsersModel.FindAll(l.ctx, "role = 0", req.Page, req.PageSize)
	GetAllApplicantResp, err := l.svcCtx.UsersModel.FindAllByPage(l.ctx, req.Page, req.PageSize, "role=0")
	// GetAllApplicantResp, err := l.svcCtx.UsersModel.FindUserByRole(l.ctx, 0)
	if err != nil {
		return nil, response.Error(500, err.Error())
	}

	var result []types.User // 将models.Users转换成types.User
	if len(GetAllApplicantResp) > 0 {
		for _, usersmodel := range GetAllApplicantResp {
			var typeUser types.User
			_ = copier.Copy(&typeUser, usersmodel)
			result = append(result, typeUser)
		}
	}
	return &types.GetAllApplicantResp{
		List: result,
	}, nil
}
