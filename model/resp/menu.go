package resp

type MenuListVO struct {
	ID        int          `gorm:"comment:id" json:"id"`
	Name      string       `gorm:"type:varchar(20);comment:菜单名" json:"name"`
	Icon      string       `gorm:"type:varchar(50);comment:菜单图标" json:"icon"`
	ParentId  int          `gorm:"comment:父菜单id" json:"parent_id"`
	Path      string       `gorm:"type:varchar(50);comment:菜单路径" json:"path"`
	Component string       `gorm:"type:varchar(50);comment:组件" json:"component"`
	Sort      int          `gorm:"type:tinyint;default:0;comment:显示排序" json:"sort"`
	Status    string       `gorm:"type:varchar(20);default:1s;comment:是否隐藏(0-否 1-是)" json:"status"`
	Children  []MenuListVO `json:"children"`
}
