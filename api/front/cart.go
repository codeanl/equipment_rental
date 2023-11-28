package front

import (
	"github.com/gin-gonic/gin"
	"outdoor_rental/api"
	"outdoor_rental/model/req"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Cart struct {
}

// SaveOrUpdateCart 添加｜｜更新菜单
func (*Cart) SaveOrUpdateCart(c *gin.Context) {
	r.SendCode(c, api.CartService.SaveOrUpdateCart(utils.BindJson[req.SaveOrUpdateCart](c)))
}

// DeleteCart 删除
func (*Cart) DeleteCart(c *gin.Context) {
	r.SendCode(c, api.CartService.DeleteCart(utils.BindJson[req.Delete](c)))
}

// GetCartList 列表
func (*Cart) GetCartList(c *gin.Context) {
	r.SuccessData(c, api.CartService.GetCartList(utils.BindQuery[req.CartList](c)))
}
