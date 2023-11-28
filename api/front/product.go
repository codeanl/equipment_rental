package front

import (
	"github.com/gin-gonic/gin"
	"outdoor_rental/api"
	"outdoor_rental/model/req"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Product struct {
}

//ProductList 列表
func (*Product) ProductList(c *gin.Context) {
	r.SuccessData(c, api.ProductService.ProductList(utils.BindQuery[req.ProductList](c)))
}

//ProductInfo 列表
func (*Product) ProductInfo(c *gin.Context) {
	r.SuccessData(c, api.ProductService.ProductInfo(utils.BindQuery[req.ProductInfo](c)))
}
