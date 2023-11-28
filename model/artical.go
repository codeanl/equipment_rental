package model

type Article struct {
	Public   `mapstructure:",squash"`
	MemberID int    `gorm:"type:tinyint;default:0;comment:用户id" json:"member_id"`
	Title    string `gorm:"type:varchar(20);comment:标题" json:"title"`
	Content  string `gorm:"type:text;comment:内容" json:"content"`
	Pic      string `gorm:"type:varchar(255);comment:封面" json:"pic"`
}
