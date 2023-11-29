package resp

type CartListVO struct {
	ID       int    ` json:"id" mapstructure:"id"`
	MemberID int    ` json:"member_id"`
	SkuID    int    ` json:"sku_id"`
	Count    int    ` json:"count"`
	Selected string `json:"selected"`

	//
	Name      string  `json:"name" `
	Pic       string  `json:"pic" `
	Desc      string  `json:"desc" `
	Price     float64 `json:"price"`
	Tag       string  `json:"tag"`
	ProductID int     `json:"product_id"`
}
