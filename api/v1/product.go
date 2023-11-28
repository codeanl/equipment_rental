package v1

import (
	"github.com/gin-gonic/gin"
	"outdoor_rental/api"
	"outdoor_rental/model/req"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
	"strconv"
)

type Product struct{}

// SaveOrUpdate 添加｜｜更新菜单
func (*Product) SaveOrUpdate(c *gin.Context) {
	r.SendCode(c, api.ProductService.SaveOrUpdate(utils.BindValidJson[req.SaveOrUpdateProduct](c)))
}

// ProductDelete 删除
func (*Product) ProductDelete(c *gin.Context) {
	r.SendCode(c, api.ProductService.ProductDelete(utils.BindJson[req.Delete](c)))
}

// ProductList 列表
func (*Product) ProductList(c *gin.Context) {
	r.SuccessData(c, api.ProductService.ProductList(utils.BindQuery[req.ProductList](c)))
}

// SkuUpdate 更新sku
func (*Product) SkuUpdate(c *gin.Context) {
	r.SendCode(c, api.ProductService.SkuUpdate(utils.BindValidJson[req.UpdateSku](c)))
}

// SkuList sku列表
func (*Product) SkuList(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Query("pid"))
	r.SuccessData(c, api.ProductService.SkuList(pid))
}
