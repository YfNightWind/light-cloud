package model

type UserInfo struct {
	Id       int
	Identity string
	Name     string
	Password string
	Email    string
}

func (table UserInfo) TableName() string {
	// 返回数据库的用户表名
	return "user_info"
}
