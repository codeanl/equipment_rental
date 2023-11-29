package service

import (
	"outdoor_rental/dao"
	"outdoor_rental/model"
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
	"time"
)

type Order struct{}

func (*Order) OrderInfo(req req.OrderInfo) resp.OrderListVO {
	order := dao.GetOne(model.Order{}, "id", req.ID)
	data := utils.CopyProperties[resp.OrderListVO](order)
	orderSku := dao.List([]model.OrderSku{}, "*", "", "order_id =  ?", data.ID)
	var dd []int
	for _, i3 := range orderSku {
		dd = append(dd, i3.SkuId)
	}
	data.Skus = dao.List([]model.ProductSku{}, "*", "", "id in ?", dd)
	return data
}

//OrderList 订单列表
func (*Order) OrderList(req req.OrderList) resp.PageResult[[]resp.OrderListVO] {
	list, count := orderDao.OrderList(req)
	//todo 使用关联取查询
	for index, i := range list {
		orderSku := dao.List([]model.OrderSku{}, "*", "", "order_id =  ?", i.ID)
		var dd []int
		for _, i3 := range orderSku {
			dd = append(dd, i3.SkuId)
		}
		list[index].Skus = dao.List([]model.ProductSku{}, "*", "", "id in ?", dd)
	}
	return resp.PageResult[[]resp.OrderListVO]{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		Total:    count,
		List:     list,
	}
}

// OrderUpdate 更新
func (*Order) OrderUpdate(req req.OrderUpdate) (code int) {
	order := model.Order{
		ID:     req.ID,
		Status: req.Status,
	}
	if req.Status == "1" {
		order.PaymentTime = time.Now()
	} else if req.Status == "3" {
		order.PickUpTime = time.Now()
	} else if req.Status == "4" {
		order.ReturnTime = time.Now()
	}
	dao.Update(&order)
	return r.OK
}
func (*Order) AddOrder(req req.AddOrder) (code int) {
	data := utils.CopyProperties[model.Order](req)
	dao.Create(&data)
	for _, sku := range req.Sku {
		sku.OrderId = data.ID
		s := utils.CopyProperties[model.OrderSku](sku)
		dao.Create(&s)
	}
	return r.OK
}

// OrderDelete 删除菜单
func (*Order) OrderDelete(req req.Delete) (code int) {
	for _, i := range req.ID {
		// 检查要删除的菜单是否存在
		existMenuById := dao.GetOne(model.Order{}, "id", i)
		if existMenuById.ID == 0 {
			return r.ERROR_CATE_NOT_EXIST
		}
	}
	dao.Delete(model.Order{}, "id in (?)", req.ID)
	dao.Delete(model.OrderSku{}, "order_id in (?)", req.ID)
	return r.OK
}
