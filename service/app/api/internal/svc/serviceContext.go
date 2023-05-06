package svc

import (
	"management-system/service/app/api/internal/config"
	"management-system/service/app/models"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config               config.Config
	RedisClient          *redis.Redis
	UsersModel           models.UsersModel
	UsersAuthModel       models.UsersAuthModel
	RolesChangeModel     models.RolesChangeModel
	RescueInfoModel      models.RescueInfoModel
	ApplicationFormsModel models.ApplicationFormsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		UsersModel:       models.NewUsersModel(sqlx.NewMysql(c.DB.DataSource)),
		UsersAuthModel:   models.NewUsersAuthModel(sqlx.NewMysql(c.DB.DataSource)),
		RolesChangeModel: models.NewRolesChangeModel(sqlx.NewMysql(c.DB.DataSource)),
		RescueInfoModel:  models.NewRescueInfoModel(sqlx.NewMysql(c.DB.DataSource)),
		ApplicationFormsModel: models.NewApplicationFormsModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
