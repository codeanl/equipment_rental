package resp

import (
	"outdoor_rental/model"
	"time"
)

// 登录 VO
type LoginVO struct {
	//相关信息
	ID       int    `json:"id"`
	Username string `json:"username"`
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

// 用户信息 VO
type UserInfoVO struct {
	ID       int          `json:"id"`
	Username string       `json:"username"`
	Nickname string       `json:"nickname"`
	Email    string       `json:"email"`
	Phone    string       `json:"phone"`
	Avatar   string       `json:"avatar"`
	Intro    string       `json:"intro"`
	Roles    []model.Role `json:"roles"`
	//Roles    []string     `json:"roles"`
	Menus []MenuListVO `json:"menus"`
}

// 用户列表 VO
//gorm:"many2many:user_role;foreignKey:ID;joinForeignKey:UserId;"：这是 GORM 的标签，它定义了与数据库表和关系有关的信息。
//many2many:user_role：指示这是一个多对多关系，它告诉 GORM 在 user 表和 role 表之间使用 user_role 表来管理这个关系。
//foreignKey:ID：表示在 user 表中，ID 列将被用作外键，用于建立与 user_role 表的关联。
//joinForeignKey:UserId：表示在 user_role 表中，UserId 列将被用作外键，用于建立与 user 表的关联。

type UserListVO struct {
	ID            int          `json:"id"`
	Avatar        string       `json:"avatar"`
	Nickname      string       `json:"nickname"`
	Username      string       `json:"username"`
	Email         string       `json:"email"`
	Phone         string       `json:"phone"`
	Status        string       `json:"status"`
	Roles         []model.Role `json:"roles" gorm:"many2many:user_role;foreignKey:ID;joinForeignKey:UserId;"`
	IpAddress     string       `json:"ip_address"`
	IpSource      string       `json:"ip_source"`
	CreatedAt     time.Time    `json:"created_at"`
	LastLoginTime time.Time    `json:"last_login_time"`
	Intro         string       `json:"intro"`
}
