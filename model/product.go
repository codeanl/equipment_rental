package model

type Product struct {
	Public
	CategoryID int64   `json:"category_id" gorm:"type:bigint;comment:商品分类id;not null"`
	Name       string  `json:"name" gorm:"type:varchar(191);comment:商品名称;not null"`
	Pic        string  `json:"pic" gorm:"type:varchar(191);comment:封面图片;not null"`
	Desc       string  `json:"desc" gorm:"type:varchar(191);comment:商品描述;not null"`
	Price      float64 `json:"price" gorm:"type: decimal(10, 2);comment:价格;not null"`
	Sale       int64   `json:"sale" gorm:"type:bigint;comment:销量;not null"`
}
