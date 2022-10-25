package model

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var Engine = Init()

// Init 初始化 Xorm创建Engine，连接数据库
func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "数据库用户:密码@(地址x.x.x.x)/数据库名称?charset=utf8")

	if err != nil {
		log.Printf("Xorm新建引擎的时候出现了一些问题：%v\n", err)
		return nil
	}

	return engine
}
