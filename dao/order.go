package dao

import (
	"fmt"
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
)

type Order struct{}

func (*Order) OrderList(req req.OrderList) (list []resp.OrderListVO, total int64) {
	db := DB.Model(&Order{}).Order("created_at DESC").Preload("Skus")
	//.Preload("OrderSku")
	if req.MinPrice != 0 {
		db = db.Where("total_amount >= ?", req.MinPrice)
	}
	if req.MaxPrice != 0 {
		db = db.Where("total_amount <= ?", req.MaxPrice)
	}
	if req.MemberId != 0 {
		db = db.Where("member_id = ?", req.MemberId)
	}
	if req.OrderType != "" {
		db = db.Where("order_type = ?", req.OrderType)
	}
	if req.PayType != "" {
		db = db.Where("pay_type = ?", req.PayType)
	}
	if req.Status != "" && req.Status != "0" {
		db = db.Where("status = ?", req.Status)
	}
	if req.Address != "" {
		db = db.Where("address LIKE ?", fmt.Sprintf("%%%s%%", req.Address))
	}
	err := db.Count(&total).Error
	if err != nil {
		return list, total
	}
	if req.PageNum > 0 && req.PageSize > 0 {
		err = db.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	return list, total
}
