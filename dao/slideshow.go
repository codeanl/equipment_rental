package dao

import (
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
)

type SlideShow struct{}

func (*SlideShow) SlideShowList(req req.SlideshowList) (list []resp.SlideshowListVO, total int64) {
	db := DB.Model(&SlideShow{})
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}
	if req.Name != "" {
		db = db.Where("name = ?", req.Name)
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
