package svc

import (
	"ShortURL/internal/config"
	"ShortURL/internal/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.ShortUrlsModel
	Redis  *redis.Redis // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewShortUrlsModel(sqlx.NewMysql(c.DataSource)),
		Redis:  redis.MustNewRedis(c.Redis),
	}
}
