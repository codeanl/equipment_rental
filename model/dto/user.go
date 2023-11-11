package dto

import "outdoor_rental/model/resp"

type UserDetailDTO struct {
	resp.LoginVO
	Password string   `json:"password"`
	Status   string   `json:"status"`
	Browser  string   `json:"browser"` //浏览器
	OS       string   `json:"os"`
	Roles    []string `json:"roles"` //角色
}

// Session 信息: 记录用户详细信息 + 是否被强退
type SessionInfo struct {
	UserDetailDTO
	IsOffline int `json:"is_offline"`
}
