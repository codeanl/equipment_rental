package model

import (
	"time"
)

type User struct {
	Public
	Username string `gorm:"type:varchar(30);comment:用户名" json:"username"`
	Password string `gorm:"type:varchar(255);comment:密码" json:"password"`
	Nickname string `gorm:"type:varchar(30);not null;comment:昵称" json:"nickname"`
	Email    string `gorm:"type:varchar(30);comment:邮箱" json:"email"`
	Phone    string `gorm:"type:varchar(255);comment:手机号" json:"phone"`
	Avatar   string `gorm:"type:varchar(1024);not null;comment:头像地址" json:"avatar"`
	Intro    string `gorm:"type:varchar(255);comment:个人简介" json:"intro"`
	Status   string `gorm:"type:varchar(30);comment:是否禁用(0-禁用 1-正常);default:'1';" json:"status"`

	IpAddress     string    `gorm:"type:varchar(20);comment:登录IP地址" json:"ip_address"`
	IpSource      string    `gorm:"type:varchar(50);comment:IP来源" json:"ip_source"`
	LastLoginTime time.Time `gorm:"comment:上次登录时间;null;" json:"last_login_time"`
}
