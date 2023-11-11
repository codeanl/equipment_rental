package model

// 分类
type Category struct {
	Public   `mapstructure:",squash"`
	Name     string `gorm:"type:varchar(20);comment:分类名" json:"name"`
	Pic      string `gorm:"type:varchar(255);comment:分类图片" json:"pic"`
	ParentId int    `gorm:"type:tinyint;default:0;comment:父菜单id" json:"parent_id"`
	Sort     int    `gorm:"type:tinyint;default:0;comment:显示排序" json:"sort"`
	Status   string `gorm:"type:varchar(50);default:0;comment:是否隐藏(0-否 1-是)" json:"status"`
}
