package req

//登录
type Login struct {
	Username string `json:"username" validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required" label:"密码"`
}

//注册
type UserAdd struct {
	Username string `json:"username" validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required" label:"密码"`
	Phone    string `json:"phone" validate:"required" label:"手机号"`
	//Email    string `json:"email" validate:"required" label:"邮箱"`
	//Code     string `json:"code" validate:"required" label:"昵称"`
}

//用户列表
type UserList struct {
	PageSize int    `form:"page_size"`
	PageNum  int    `form:"page_num"`
	Nickname string `form:"nickname"`
	Username string `form:"username"`
	Email    string `form:"email"`
	Phone    string `form:"phone"`
	Status   string `form:"status"`
}

//更新用户
type UpdateUser struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `gorm:"type:varchar(30);comment:邮箱" json:"email"`
	Phone    string `gorm:"type:varchar(30);comment:手机号" json:"phone"`
	Avatar   string `gorm:"type:varchar(1024);not null;comment:头像地址" json:"avatar"`
	Intro    string `gorm:"type:varchar(255);comment:个人简介" json:"intro"`
	Status   string `gorm:"type:varchar(30);comment:是否禁用(0-禁用 1-正常)" json:"status"`
	RoleIds  []int  `json:"role_ids"`
}

//更新密码
type SetPass struct {
	OldPass     string `json:"old_pass"`
	NewPass     string `json:"new_pass"`
	OnceNewPass string `json:"once_new_pass"`
}

//更新个人信息
type UserUpdateInfo struct {
	Nickname string `json:"nickname"`
	Email    string `gorm:"type:varchar(30);comment:邮箱" json:"email"`
	Phone    string `gorm:"type:varchar(30);comment:手机号" json:"phone"`
	Avatar   string `gorm:"type:varchar(1024);not null;comment:头像地址" json:"avatar"`
	Intro    string `gorm:"type:varchar(255);comment:个人简介" json:"intro"`
}
