package v1

import (
	"github.com/gin-gonic/gin"
	"outdoor_rental/model/req"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
	"strconv"
)

type Product struct{}

// SaveOrUpdate 添加｜｜更新菜单
func (*Product) SaveOrUpdate(c *gin.Context) {
	r.SendCode(c, productService.SaveOrUpdate(utils.BindValidJson[req.SaveOrUpdateProduct](c)))
}

// ProductDelete 删除
func (*Product) ProductDelete(c *gin.Context) {
	r.SendCode(c, productService.ProductDelete(utils.BindJson[req.Delete](c)))
}

// ProductList 用户列表
func (*Product) ProductList(c *gin.Context) {
	r.SuccessData(c, productService.ProductList(utils.BindQuery[req.ProductList](c)))
}

// SkuUpdate 更新sku
func (*Product) SkuUpdate(c *gin.Context) {
	r.SendCode(c, productService.SkuUpdate(utils.BindValidJson[req.UpdateSku](c)))
}

// SkuList sku列表
func (*Product) SkuList(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Query("pid"))
	r.SuccessData(c, productService.SkuList(pid))
}
