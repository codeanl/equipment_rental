package model

type Cart struct {
	Public   `mapstructure:",squash"`
	MemberID int  `gorm:"comment:MemberID" json:"member_id"`
	SkuID    int  `gorm:"comment:skuId" json:"sku_id"`
	Count    int  `gorm:"comment:数量" json:"count"`
	Selected bool `gorm:"comment:数量" json:"selected"`
}
