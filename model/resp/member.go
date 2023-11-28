package resp

import (
	"time"
)

type MemberListVO struct {
	ID        int    `json:"id"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	IpAddress string `json:"ip_address"`
	IpSource  string `json:"ip_source"`

	CreatedAt     time.Time `json:"created_at"`
	LastLoginTime time.Time `json:"last_login_time"`
	Intro         string    `json:"intro"`
}

// 用户信息 VO
type MemberInfoVO struct {
	ID       int    `json:"id"`
	Phone    string `json:"phone"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Intro    string `json:"intro"`

	IpAddress     string    `gorm:"type:varchar(20);comment:登录IP地址" json:"ip_address"`
	IpSource      string    `gorm:"type:varchar(50);comment:IP来源" json:"ip_source"`
	LastLoginTime time.Time `gorm:"comment:上次登录时间;null;" json:"last_login_time"`
}

type LoginMemberVO struct {
	//相关信息
	ID       int    `json:"id"`
	Phone    string `json:"phone"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Intro    string `json:"intro"`
	//登录信息
	IpAddress     string    `json:"ip_address"`
	IpSource      string    `json:"ip_source"`
	LastLoginTime time.Time `json:"last_login_time"`
	//密钥
	Token string `json:"token"`
}
