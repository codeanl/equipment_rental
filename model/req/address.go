package req

type AddressList struct {
	PageSize int `form:"page_size"`
	PageNum  int `form:"page_num"`
	MemberId int `form:"member_id"`
}
type SaveOrUpdateAddress struct {
	ID        int    `json:"id"`
	MemberID  int    `json:"member_id"`
	Name      string ` json:"name"`
	Phone     string `json:"phone"`
	Address   string ` json:"address"`
	IsDefault string `json:"is_default"`
}
