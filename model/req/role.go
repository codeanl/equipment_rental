package req

// SaveOrUpdateRole 新增||编辑 角色, 关联维护 role_resource, role_menu
type SaveOrUpdateRole struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Label  string `json:"label"`
	Status string `json:"status"`
	//ApiIds  []int  `json:"api_ids"`
	MenuIds []int `json:"menu_ids"` //*
}

type RoleList struct {
	PageSize int    `form:"page_size"`
	PageNum  int    `form:"page_num"`
	Name     string `form:"name"`
	Label    string `form:"label"`
	Status   string `form:"status"`
}
