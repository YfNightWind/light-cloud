package model

import "time"

type UserInfo struct {
	Id        int
	Identity  string
	Name      string
	Password  string
	Email     string
	Avatar    string
	Capacity  int
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (table UserInfo) TableName() string {
	// 返回数据库的用户表名
	return "user_info"
}
