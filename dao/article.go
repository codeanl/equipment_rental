package dao

import (
	"fmt"
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
)

type Article struct{}

func (*Article) ArticleList(req req.ArticleList) (list []resp.ArticleListVO, total int64) {
	sort := "RAND()"
	if req.SorType == 1 {
		sort = "created_at DESC"
	}
	if req.SorType == 2 {
		sort = "created_at DESC"
	}
	db := DB.Model(&Article{}).Order(sort)
	//db := DB.Model(&Article{}).Order("created_at DESC")
	if req.MemberId != 0 {
		db = db.Where("member_id = ?", req.MemberId)
	}
	if req.Title != "" {
		db = db.Where("title LIKE ?", fmt.Sprintf("%%%s%%", req.Title))
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
