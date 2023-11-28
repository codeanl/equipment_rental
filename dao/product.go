package dao

import (
	"fmt"
	"outdoor_rental/model"
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
)

type Product struct{}

func (*Product) ProductList(req req.ProductList) (list []resp.ProductListVO, total int64) {
	sort := "RAND()"
	if req.SorType == 1 {
		sort = "created_at DESC"
	}
	if req.SorType == 2 {
		sort = "price ASC"
	}
	if req.SorType == 3 {
		sort = "price DESC"
	}
	db := DB.Model(&Product{}).Order(sort).Preload("ProductImg")
	//db := m.conn.Model(&Product{}).Order("created_at DESC")
	if req.CategoryId != 0 {
		var count int64
		err := DB.Model(&model.Category{}).Where("parent_id = ?", req.CategoryId).Count(&count).Error
		if err != nil {
			return list, 0
		}
		// 检查是否有父级分类
		if count == 0 {
			db = db.Where("category_id = ?", req.CategoryId)
		} else {
			var category []model.Category
			err := DB.Model(&model.Category{}).Where("parent_id = ?", req.CategoryId).Find(&category).Error
			if err != nil {
				return list, 0
			}
			var ids []int64
			for _, i := range category {
				ids = append(ids, int64(i.ID))
			}
			db = db.Where("category_id IN (?)", ids)
		}
	}
	if req.MinPrice != 0 {
		db = db.Where("price >= ?", req.MinPrice)
	}
	if req.MaxPrice != 0 {
		db = db.Where("price <= ?", req.MaxPrice)
	}
	if req.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Name))
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

func (*Product) GetProductSkuTag(id int) (labels []string) {
	DB.Model(model.ProductSku{}).Where("product_id=?", id).Pluck("tag", &labels)
	return
}

func (*Product) ProductInfo(id int) (resp resp.ProductListVO) {
	DB.Model(model.Product{}).Where("id=?", id).Preload("Skus").Preload("ProductImg").Preload("ProductDetailImg").First(&resp)
	return resp
}
