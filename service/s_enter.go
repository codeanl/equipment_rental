package service

import "outdoor_rental/dao"

const (
	KEY_CODE   = "code:"   // 验证码
	KEY_USER   = "user:"   // 记录用户
	KEY_DELETE = "delete:" //? 记录强制下线用户?

)

var (
	userDao    dao.User
	roleDao    dao.Role
	menuDao    dao.Menu
	productDao dao.Product
	memberDao  dao.Member
	orderDao   dao.Order
)
