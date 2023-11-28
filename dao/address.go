package dao

import (
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
)

type Address struct{}

func (*Address) AddressList(req req.AddressList) (list []resp.AddressListVO, total int64) {
	db := DB.Model(&Address{}).Order("created_at DESC")
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
