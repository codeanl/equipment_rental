package resp

import (
	"outdoor_rental/model"
)

type ProductListVO struct {
	ID         int                `json:"id"`
	CategoryID int64              `json:"category_id"`
	Name       string             `json:"name"`
	Pic        string             `json:"pic"`
	Desc       string             `json:"desc"`
	Sale       int64              `json:"sale"`
	ProductImg []model.ProductImg `json:"imgs" gorm:"foreignKey:product_id"`
}
type SkuListVO struct {
	ID        int     `json:"id"`
	ProductID int64   `json:"product_id"`
	Name      string  `json:"name"`
	Pic       string  `json:"pic"`
	Desc      string  `json:"desc"`
	Price     float64 `json:"price"`
	Stock     int64   `json:"stock"`
	Sale      int64   `json:"sale"`
	Tag       string  `json:"tag"`
}
