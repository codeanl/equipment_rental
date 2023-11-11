package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"outdoor_rental/middleware"
)

func AdminRouter() http.Handler {
	gin.SetMode("debug")
	r := gin.New()
	r.Use(middleware.Logger()) // 自定义的 zap 日志中间件
	r.Use(middleware.Cors())   // 跨域中间件

	base := r.Group("api")
	{
		base.POST("/login", userAPI.Login) //登录
	}
	base.Use(middleware.JWTAuth()) // JWT 鉴权中间件
	//用户模块
	user := base.Group("/user")
	{
		user.GET("/info", userAPI.Profile)               //个人信息
		user.POST("", userAPI.UserAdd)                   //添加用户
		user.PUT("", userAPI.UserUpdate)                 // 更新用户
		user.POST("del", userAPI.UserDelete)             // 删除用户
		user.GET("", userAPI.UserInfo)                   //用户信息
		user.GET("/list", userAPI.UserList)              // 用户列表
		user.POST("/setPass", userAPI.UserSetPass)       // 修改密码
		user.POST("/updateInfo", userAPI.UserUpdateInfo) // 修改个人信息 //TODO 更换邮箱/手机号需要接受验证码
	}
	//角色模块
	role := base.Group("/role")
	{
		role.GET("/list", roleAPI.GetTreeList) // 角色列表(树形)
		role.POST("", roleAPI.SaveOrUpdate)    // 新增/编辑菜单
		role.POST("/del", roleAPI.Delete)      // 删除角色
	}
	//菜单模块
	menu := base.Group("/menu")
	{
		menu.GET("/list", menuAPI.GetTreeList)     // 树形菜单列
		menu.POST("", menuAPI.SaveOrUpdate)        // 新增/编辑菜单  //TODO pId修改不能变成0
		menu.POST("del", menuAPI.Delete)           // 删除菜单
		menu.GET("/userMenu", menuAPI.GetUserMenu) // 获取当前用户的菜单
	}
	//分类模块
	category := base.Group("/category")
	{
		category.GET("/list", categoryAPI.GetTreeList) // 列表
		category.POST("", categoryAPI.SaveOrUpdate)    // 新增/编辑 //TODO pId修改不能变成0
		category.POST("del", categoryAPI.Delete)       // 删除
	}
	//商品模块
	product := base.Group("/product")
	{
		product.GET("/list", productAPI.ProductList)   // 列表
		product.POST("", productAPI.SaveOrUpdate)      // 新增/编辑 TODO pId修改不能变成0
		product.POST("/del", productAPI.ProductDelete) // 删除
	}
	sku := base.Group("/sku")
	{
		sku.GET("/list", productAPI.SkuList) // 列表
		sku.POST("", productAPI.SkuUpdate)   // 编辑
	}
	//会员模块
	member := base.Group("/member")
	{
		member.GET("", memberAPI.MemberInfo)        //会员信息
		member.GET("/list", memberAPI.MemberList)   // 会员列表
		member.PUT("", memberAPI.MemberUpdate)      // 更新会员
		member.POST("/del", memberAPI.MemberDelete) // 删除会员
	}
	//订单模块
	order := base.Group("/order")
	{
		//order.GET("", orderAPI.OrderInfo)      //会员信息
		order.GET("/list", orderAPI.OrderList) // 订单列表
		order.PUT("", orderAPI.OrderUpdate)    // 更新订单
		order.DELETE("", orderAPI.OrderDelete) // 删除订单
	}
	return r
}
