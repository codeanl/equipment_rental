package resp

import "outdoor_rental/model"

type ArticleListVO struct {
	ID       int          ` json:"id" mapstructure:"id"`
	MemberID int          ` json:"member_id"`
	Title    string       ` json:"title"`
	Content  string       ` json:"content"`
	Pic      string       ` json:"pic"`
	Member   model.Member `json:"member"`
}
