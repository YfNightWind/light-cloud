package svc

import (
	"github.com/go-redis/redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"light-cloud/src/core/internal/config"
	"light-cloud/src/core/internal/middleware"
	"light-cloud/src/model"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	SQL    *xorm.Engine
	RDB    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		SQL:    model.Init(c.MySQL.DataSourceName),
		RDB:    model.InitRedis(c),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
