package front

import (
	"github.com/gin-gonic/gin"
	"outdoor_rental/api"
	"outdoor_rental/model/req"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Order struct{}

func (*Order) AddOrder(c *gin.Context) {
	r.SendCode(c, api.OrderService.AddOrder(utils.BindJson[req.AddOrder](c)))
}

func (*Order) UpdateOrder(c *gin.Context) {
	r.SendCode(c, api.OrderService.OrderUpdate(utils.BindJson[req.OrderUpdate](c)))
}

func (*Order) DeleteOrder(c *gin.Context) {
	r.SendCode(c, api.OrderService.OrderDelete(utils.BindJson[req.Delete](c)))
}

func (*Order) GetOrderList(c *gin.Context) {
	r.SuccessData(c, api.OrderService.OrderList(utils.BindQuery[req.OrderList](c)))
}
func (*Order) OrderInfo(c *gin.Context) {
	r.SuccessData(c, api.OrderService.OrderInfo(utils.BindQuery[req.OrderInfo](c)))
}
