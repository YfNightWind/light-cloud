package model

type UserBasic struct {
	Id       int
	Identity string
	Name     string
	Password string
	Email    string
}

func (table UserBasic) TableName() string {
	// 返回数据库的用户表名
	return "user_basic"
}
