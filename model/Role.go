package model

// 角色
type Role struct {
	Public
	Name   string `gorm:"type:varchar(20);comment:角色名" json:"name"`
	Label  string `gorm:"type:varchar(50);comment:角色描述" json:"label"`
	Status string `gorm:"type:varchar(50);comment:是否禁用(0-否 1-是);" json:"status"`
}
