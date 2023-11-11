package model

// 用户-角色 关联
type UserRole struct {
	UserId int `json:"user_id"`
	RoleId int `json:"role_id"`
}
