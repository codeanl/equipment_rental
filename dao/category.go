package dao

import (
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
)

type Category struct{}

func (*Category) CategoryList(req req.FrontCategoryList) (list []resp.FrontCategoryListVO, total int64) {
	db := DB.Model(&Category{}).Order("created_at DESC")
	if req.ParentId == 0 {
		db = db.Where("parent_id = ?", 0)
	}
	if req.ParentId != 0 {
		db = db.Where("parent_id = ?", req.ParentId)
	}
	err := db.Count(&total).Error
	if err != nil {
		return list, total
	}

	err = db.Find(&list).Error
	return list, total
}
