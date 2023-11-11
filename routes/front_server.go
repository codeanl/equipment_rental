package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"outdoor_rental/middleware"
)

func FrontRouter() http.Handler {
	gin.SetMode("debug")
	r := gin.New()
	r.Use(middleware.Logger()) // 自定义的 zap 日志中间件
	r.Use(middleware.Cors())   // 跨域中间件
	//base := r.Group("front")
	{
		//base.POST("/login", memberFrontAPI.Login) //登录
	}
	return r
}
