package dao

import (
	"outdoor_rental/model"
	"outdoor_rental/model/req"
)

type Cart struct{}

func (*Cart) CartList(req req.CartList) (list []model.Cart, total int64) {
	//func (*Cart) CartList(req req.CartList) (list []resp.CartListVO, total int64) {
	db := DB.Model(&Cart{}).Order("created_at DESC")
	//.Preload("OrderSku")
	if req.MemberId != 0 {
		db = db.Where("member_id = ?", req.MemberId)
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
