package dao

import (
	"outdoor_rental/model"
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
)

type Role struct{}

// GetRolesByUserId 根据用户id查询用户角色
func (*Role) GetRolesByUserId(userId int) (roles []string) {
	DB.Table("role , user_role ").
		Where("role.id = user_role.role_id AND user_role.user_id = ?", userId).
		Pluck("label", &roles) // 将单列查询为切片
	return
}

// GetRolesByUserId 根据用户id查询用户角色
//func (*Role) GetRolesNameByUserId(userId int) (roles []string) {
func (*Role) GetRolesNameByUserId(userId int) (roles []model.Role) {
	DB.Table("role , user_role ").
		Where("role.id = user_role.role_id AND user_role.user_id = ?", userId).Find(&roles)
	//Where("role.id = user_role.role_id AND user_role.user_id = ?", userId).Pluck("name", &roles)
	return
}

// GetLabelsByRoleIds 根据 [角色id] 获取 [角色标签列表]
func (*Role) GetLabelsByRoleIds(id []int) (labels []string) {
	DB.Model(model.Role{}).Where("id in (?)", id).Pluck("label", &labels)
	return
}

// 获取角色列表[非树形]
func (*Role) GetList(req req.RoleList) (list []resp.RoleVo, total int64) {
	list = make([]resp.RoleVo, 0)
	db := DB.Table("role")
	// 查询条件
	if req.Name != "" {
		db = db.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Label != "" {
		db = db.Where("label like ?", "%"+req.Label+"%")
	}
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0
	}
	if req.PageNum > 0 && req.PageSize > 0 {
		err = db.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	return list, total
}

// 根据 [角色id] 查询出 [资源id列表]
//func (*Role) GetApiByRoldId(roleId int) (apiIds []int) {
//	DB.Model(&model.RoleApi{}).
//		Where("role_id = ?", roleId).
//		Pluck("api_id", &apiIds)
//	return
//}

// 根据 [角色id] 查询出 [目录id列表]
func (*Role) GetMenusByRoleId(roleId int) (menuIds []int) {
	DB.Model(&model.RoleMenu{}).
		Where("role_id = ?", roleId).
		Pluck("menu_id", &menuIds)
	return
}
