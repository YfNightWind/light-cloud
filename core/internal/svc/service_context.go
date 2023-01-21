package svc

import (
	"github.com/go-redis/redis/v9"
	"light-cloud/src/core/internal/config"
	"light-cloud/src/core/model"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: model.Init(c.MySQL.DataSourceName),
		RDB:    model.InitRedis(c),
	}
}
