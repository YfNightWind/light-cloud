package model

import (
	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
	"light-cloud/src/core/define"
	"log"
	"xorm.io/xorm"
)

// Engine 初始化数据库
var Engine = Init()

// RDB 初始化Redis
var RDB = InitRedis()

// Init 初始化 Xorm创建Engine，连接数据库
func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "数据库用户:密码@(地址x.x.x.x)/数据库名称?charset=utf8")

	if err != nil {
		log.Printf("Xorm新建引擎的时候出现了一些问题：%v\n", err)
		return nil
	}

	return engine
}

// InitRedis 初始化Redis
func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     define.RedisAddress,
		Password: define.RedisPassword,
		DB:       0, // use default DB
	})
}
