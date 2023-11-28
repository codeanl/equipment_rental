package req

//用户列表
type ProductList struct {
	PageSize   int     `form:"page_size"`
	PageNum    int     `form:"page_num"`
	Name       string  `form:"name"`
	CategoryId int64   `form:"category_id"`
	SorType    int     `form:"sort_type"`
	MaxPrice   float64 `form:"max_price"`
	MinPrice   float64 `form:"min_price"`
}

type SaveOrUpdateProduct struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	CategoryId int64    `json:"category_id"`
	Pic        string   `json:"pic"`
	Desc       string   `json:"desc"`
	Price      float64  `json:"price"`
	Img        []string `json:"img"`
	Size       []Size   `json:"size"`
}
type Size struct {
	Name     string   `json:"name"`
	SizeData []string `json:"size_data"`
}

type UpdateSku struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Pic   string  `json:"pic"`
	Desc  string  `json:"desc"`
	Price float64 `json:"price"`
	Stock int64   `json:"stock"`
	Sale  int64   `json:"sale"`
}

type ProductInfo struct {
	ID int `form:"id"`
}
