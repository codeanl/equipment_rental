package resp

type SlideshowListVO struct {
	ID     int    `json:"id"`
	Name   string `gorm:"type:varchar(20);comment:菜单名" json:"name"`
	Url    string `gorm:"type:varchar(255);comment:菜单图标" json:"url"`
	Status string `gorm:"type:varchar(50);default:0;comment:是否隐藏(0-否 1-是)" json:"status"`
}
