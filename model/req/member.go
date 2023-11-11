package req

//会员列表
type MemberList struct {
	PageSize int    `form:"page_size"`
	PageNum  int    `form:"page_num"`
	Nickname string `form:"nickname"`
	Email    string `form:"email"`
	Phone    string `form:"phone"`
	Status   string `form:"status"`
}

//更新用户
type MemberUser struct {
	ID       int    `json:"id"`
	Phone    string `gorm:"type:varchar(30);comment:手机号" json:"phone"`
	Nickname string `json:"nickname"`
	Email    string `gorm:"type:varchar(30);comment:邮箱" json:"email"`
	Avatar   string `gorm:"type:varchar(1024);not null;comment:头像地址" json:"avatar"`
	Intro    string `gorm:"type:varchar(255);comment:个人简介" json:"intro"`
	Status   string `gorm:"type:varchar(30);comment:是否禁用(0-禁用 1-正常)" json:"status"`
}
