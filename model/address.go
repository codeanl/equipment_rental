package model

type Address struct {
	Public    `mapstructure:",squash"`
	MemberID  int    `gorm:"comment:MemberID" json:"member_id"`
	Name      string `gorm:"comment:姓名" json:"name"`
	Phone     string `gorm:"comment:手机号" json:"phone"`
	Address   string `gorm:"comment:地址" json:"address"`
	IsDefault string `gorm:"comment:是否默认" json:"is_default"`
}
