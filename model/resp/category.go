package resp

type CategoryListVO struct {
	ID       int              `gorm:"comment:id" json:"id"`
	Name     string           `gorm:"type:varchar(20);comment:分类名" json:"name"`
	Pic      string           `gorm:"type:varchar(50);comment:分类图标" json:"pic"`
	ParentId int              `gorm:"comment:父菜单id" json:"parent_id"`
	Sort     int              `gorm:"type:tinyint;default:0;comment:显示排序" json:"sort"`
	Status   string           `gorm:"type:varchar(20);default:1s;comment:是否隐藏(0-否 1-是)" json:"status"`
	Children []CategoryListVO `json:"children"`
}

type FrontCategoryListVO struct {
	ID       int    `gorm:"comment:id" json:"id"`
	Name     string `gorm:"type:varchar(20);comment:分类名" json:"name"`
	Pic      string `gorm:"type:varchar(50);comment:分类图标" json:"pic"`
	ParentId int    `gorm:"comment:父菜单id" json:"parent_id"`
	Sort     int    `gorm:"type:tinyint;default:0;comment:显示排序" json:"sort"`
	Status   string `gorm:"type:varchar(20);default:1s;comment:是否隐藏(0-否 1-是)" json:"status"`
}

type GetListNextCateAndSpu struct {
	ID       int             `gorm:"comment:id" json:"id"`
	Name     string          `gorm:"type:varchar(20);comment:分类名" json:"name"`
	Pic      string          `gorm:"type:varchar(50);comment:分类图标" json:"pic"`
	ParentId int             `gorm:"comment:父菜单id" json:"parent_id"`
	Sort     int             `gorm:"type:tinyint;default:0;comment:显示排序" json:"sort"`
	Status   string          `gorm:"type:varchar(20);default:1s;comment:是否隐藏(0-否 1-是)" json:"status"`
	Product  []ProductListVO `json:"product"`
}
