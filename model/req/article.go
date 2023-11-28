package req

type ArticleList struct {
	PageSize int    `form:"page_size"`
	PageNum  int    `form:"page_num"`
	MemberId int    `form:"member_id"`
	Title    string `form:"title"`
	SorType  int    `form:"sort_type"`
}
type SaveOrUpdateArticle struct {
	ID       int    `json:"id"`
	MemberID int    `json:"member_id"`
	Title    string ` json:"title"`
	Content  string `json:"content"`
	Pic      string ` json:"pic"`
}
