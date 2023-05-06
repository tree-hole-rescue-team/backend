package rescue

import (
	"context"
	"encoding/json"
	"strconv"

	"management-system/common/response"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"

	red "github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
)

type DestroyTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDestroyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DestroyTokenLogic {
	return &DestroyTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DestroyTokenLogic) DestroyToken(req *types.DestroyTokenReq) (resp *types.DestroyTokenResp, err error) {
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()
	// 先在缓存里找，有就删除
	key := "management-system token id" + strconv.Itoa(int(userId))
	token, err := l.svcCtx.RedisClient.Get(key)
	if err != nil && err != red.Nil {
		return nil, response.Error(500, "redis error:"+err.Error())
	}
	if token != "" {
		_, err = l.svcCtx.RedisClient.Del(key)
		if err != nil {
			return nil, response.Error(500, "redis error:"+err.Error())
		}
	}

	return
}
