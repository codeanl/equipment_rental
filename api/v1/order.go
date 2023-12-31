package v1

import (
	"github.com/gin-gonic/gin"
	"outdoor_rental/api"
	"outdoor_rental/model/req"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Order struct{}

// OrderList 列表
func (*Order) OrderList(c *gin.Context) {
	r.SuccessData(c, api.OrderService.OrderList(utils.BindQuery[req.OrderList](c)))
}

// OrderUpdate 更新
func (*Order) OrderUpdate(c *gin.Context) {
	r.SendCode(c, api.OrderService.OrderUpdate(utils.BindValidJson[req.OrderUpdate](c)))
}

// OrderDelete 删除
func (*Order) OrderDelete(c *gin.Context) {
	r.SendCode(c, api.OrderService.OrderDelete(utils.BindJson[req.Delete](c)))
}
