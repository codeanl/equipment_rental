package model

type ProductSku struct {
	Public
	ProductID int64   `json:"product_id" gorm:"type:bigint;comment:商品id;not null"`
	Name      string  `json:"name" gorm:"type:varchar(191);comment:sku名称;not null"`
	Pic       string  `json:"pic" gorm:"type:varchar(191);comment:封面图片;not null"`
	Desc      string  `json:"desc" gorm:"type:varchar(191);comment:商品描述;not null"`
	Price     float64 `json:"price" gorm:"type: decimal(10, 2);comment:价格;not null"`
	Stock     int64   `json:"stock" gorm:"type:bigint;comment:库存;not null;default:0"`
	Sale      int64   `json:"sale" gorm:"type:bigint;comment:销量;not null"`
	Tag       string  `json:"tag" gorm:"type:varchar(191);comment:销量;not null"`
}
