package resp

import (
	"outdoor_rental/model"
)

type ProductListVO struct {
	ID               int                      `json:"id"`
	CategoryID       int64                    `json:"category_id"`
	Name             string                   `json:"name"`
	Pic              string                   `json:"pic"`
	Desc             string                   `json:"desc"`
	Sale             int64                    `json:"sale"`
	Price            float64                  `json:"price"`
	ProductImg       []model.ProductImg       `json:"imgs" gorm:"foreignKey:product_id"`
	ProductDetailImg []model.ProductDetailImg `json:"detailImgs" gorm:"foreignKey:product_id"`
	//Skus             []SkuListVO              `json:"skus"`
	Skus []model.ProductSku `json:"skus" gorm:"foreignKey:product_id"`
	//SpecList []SpecList         `json:"spec_list"`
}

type SpecList struct {
	Name string `json:"name"`
	List []List `json:"list"`
}
type List struct {
	Name string `json:"name"`
}

type SkuListVO struct {
	ID         int      `json:"id"`
	ProductID  int64    `json:"product_id"`
	Name       string   `json:"name"`
	Pic        string   `json:"pic"`
	Desc       string   `json:"desc"`
	Price      float64  `json:"price"`
	Stock      int64    `json:"stock"`
	Sale       int64    `json:"sale"`
	Tag        string   `json:"tag"`
	SkuNameArr []string `json:"sku_name_arr"`
}
