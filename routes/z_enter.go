package routes

import v1 "outdoor_rental/api/v1"
import front "outdoor_rental/api/front"

// 后台接口
var (
	userAPI     v1.User
	roleAPI     v1.Role
	menuAPI     v1.Menu
	categoryAPI v1.Category
	productAPI  v1.Product
	memberAPI   v1.Member
	orderAPI    v1.Order
)

//前台接口
var (
	memberFrontAPI front.Member
)
