package resp

import "time"

// 角色 + 资源id列表 + 菜单id列表
type RoleVo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Label     string    `json:"label"`
	CreatedAt time.Time `json:"created_at"`
	Status    string    `json:"status"`
	//ApiIds    []int     `json:"api_ids" gorm:"-"`
	MenuIds []int `json:"menu_ids" gorm:"-"`
}
