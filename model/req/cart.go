package req

type CartList struct {
	PageSize int `form:"page_size"`
	PageNum  int `form:"page_num"`
	MemberId int `form:"member_id"`
}
type SaveOrUpdateCart struct {
	ID       int  `json:"id"`
	MemberID int  `json:"member_id"`
	SkuID    int  `json:"sku_id"`
	Count    int  `json:"count"`
	Selected bool `json:"selected"`
}
