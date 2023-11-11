package dao

import "outdoor_rental/model"

type Menu struct{}

// 根据 userInfoId 获取菜单列表(非树形结构): 关联 user_role, role_menu, menu 表
func (*Menu) GetMenusByUserInfoId(userInfoId int) []model.Menu {
	list := make([]model.Menu, 0)
	DB.Table("user_role ur").
		Distinct("m.id", "name", "path", "component", "icon", "status", "parent_id", "sort").
		Where("user_id = ?", userInfoId).
		Joins("JOIN role_menu rm ON ur.role_id = rm.role_id").
		Joins("JOIN menu m ON rm.menu_id = m.id").
		Find(&list)
	return list
}
