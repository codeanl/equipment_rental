package service

import (
	"outdoor_rental/dao"
	"outdoor_rental/model"
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Cart struct{}

func (*Cart) SaveOrUpdateCart(req req.SaveOrUpdateCart) (code int) {
	if req.ID != 0 {
		cart := model.Cart{
			Public:   model.Public{ID: req.ID},
			Count:    req.Count,
			Selected: req.Selected,
		}
		dao.Update(&cart)
	} else {
		data := utils.CopyProperties[model.Cart](req)
		dao.Create(&data)
	}
	return r.OK
}

//OrderList 订单列表
func (*Cart) GetCartList(req req.CartList) resp.PageResult[[]resp.CartListVO] {
	list, count := cartDao.CartList(req)
	data := utils.CopyProperties[[]resp.CartListVO](list)
	//todo 使用关联取查询
	for index, i := range list {
		sku := dao.GetOne(model.ProductSku{}, "id", i.SkuID)
		data[index].Name = sku.Name
		data[index].Pic = sku.Pic
		data[index].Desc = sku.Desc
		data[index].Price = sku.Price
		data[index].Tag = sku.Tag
		data[index].ProductID = int(sku.ProductID)
	}
	return resp.PageResult[[]resp.CartListVO]{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		Total:    count,
		List:     data,
	}
}

func (*Cart) DeleteCart(req req.Delete) (code int) {
	for _, i := range req.ID {
		// 检查要删除的菜单是否存在
		existMenuById := dao.GetOne(model.Cart{}, "id", i)
		if existMenuById.ID == 0 {
			return r.ERROR_CATE_NOT_EXIST
		}
	}
	// 删除菜单
	dao.Delete(model.Cart{}, "id in (?)", req.ID)
	return r.OK
}
