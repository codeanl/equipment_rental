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
	base := r.Group("/front")
	{
		base.POST("/login", memberFrontAPI.Login) //登录
	}

	index := base.Group("/index")
	{
		index.GET("/slideShowList", slideshowFrontAPI.SlideshowList) //轮播图列表
		index.GET("/productList", productFrontAPI.ProductList)       //商品列表
		index.GET("/productInfo", productFrontAPI.ProductInfo)       //商品
	}
	cate := base.Group("/category")
	{
		cate.GET("/categoryFirst", categoryFrontAPI.GetList)             //轮播图列表
		cate.GET("/categoryAll", categoryFrontAPI.GetListNextCateAndSpu) //轮播图列表
	}

	cart := base.Group("/cart")
	{
		cart.GET("/list", cartFrontAPI.GetCartList)  // 列表
		cart.POST("", cartFrontAPI.SaveOrUpdateCart) // 新增/编辑
		cart.POST("/del", cartFrontAPI.DeleteCart)   // 删除
	}
	address := base.Group("/address")
	{
		address.GET("/list", addressFrontAPI.GetAddressList)  // 列表
		address.POST("", addressFrontAPI.SaveOrUpdateAddress) // 新增/编辑
		address.POST("/del", addressFrontAPI.DeleteAddress)   // 删除
	}
	article := base.Group("/article")
	{
		article.GET("/list", articleFrontAPI.GetArticleList)  // 列表
		article.POST("", articleFrontAPI.SaveOrUpdateArticle) // 新增/编辑
		article.POST("/del", articleFrontAPI.DeleteArticle)   // 删除
	}
	order := base.Group("/order")
	{
		order.GET("/list", orderFrontAPI.GetOrderList) // 列表
		order.POST("/update", orderFrontAPI.UpdateOrder)
		order.POST("/add", orderFrontAPI.AddOrder)
		order.POST("/del", orderFrontAPI.DeleteOrder) // 删除
		order.GET("/info", orderFrontAPI.OrderInfo)   // 删除
	}
	return r
}
