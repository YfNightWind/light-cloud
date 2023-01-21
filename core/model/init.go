package model

import (
	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
	"light-cloud/src/core/internal/config"
	"log"
	"xorm.io/xorm"
)

// Init 初始化 Xorm创建Engine，连接数据库
func Init(dataSourceName string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", dataSourceName)

	if err != nil {
		log.Printf("Xorm新建引擎的时候出现了一些问题：%v\n", err)
		return nil
	}

	return engine
}

// InitRedis 初始化Redis
func InitRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Address,
		Password: c.Redis.Password,
		DB:       0, // use default DB
	})
}
