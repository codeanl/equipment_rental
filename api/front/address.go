package front

import (
	"github.com/gin-gonic/gin"
	"outdoor_rental/api"
	"outdoor_rental/model/req"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Address struct {
}

// SaveOrUpdateCart 添加｜｜更新菜单
func (*Address) SaveOrUpdateAddress(c *gin.Context) {
	r.SendCode(c, api.AddressService.SaveOrUpdateAddress(utils.BindJson[req.SaveOrUpdateAddress](c)))
}

// DeleteCart 删除
func (*Address) DeleteAddress(c *gin.Context) {
	r.SendCode(c, api.AddressService.DeleteAddress(utils.BindJson[req.Delete](c)))
}

// GetCartList 列表
func (*Address) GetAddressList(c *gin.Context) {
	r.SuccessData(c, api.AddressService.GetCartAddress(utils.BindQuery[req.AddressList](c)))
}
