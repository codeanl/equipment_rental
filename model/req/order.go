package req

import "time"

type OrderList struct {
	PageSize int `form:"page_size"`
	PageNum  int `form:"page_num"`
	MemberId int `form:"member_id"`

	OrderType string `form:"order_type"`
	PayType   string `form:"pay_type"`
	Status    string `form:"status"`
	Address   string `form:"address"`

	MaxPrice float64 `form:"max_price"`
	MinPrice float64 `form:"min_price"`
}

type OrderUpdate struct {
	ID int `gorm:"primary_key;auto_increment" json:"id" mapstructure:"id"`

	TotalAmount   float64 `json:"total_amount" gorm:"type:decimal(10, 2) ;comment:订单总金额;not null"`
	FreightAmount float64 `json:"freight_amount" gorm:"type:decimal(10, 2) ;comment:运费金额;not null"` //运费金额
	ProductAmount float64 `json:"pay_amount" gorm:"type:decimal(10, 2) ;comment:商品金额;not null"`     //商品金额
	PledgeAmount  float64 `json:"pledge_amount" gorm:"type:decimal(10, 2) ;comment:押金;not null"`    //押金

	Status      string    `json:"status" gorm:"type:varchar(191);comment:订单状态;not null"` //1->待支付 2->已支付  3->待取货 4->租赁中 5->已归还
	PaymentTime time.Time `json:"payment_time" gorm:";comment:支付时间;null"`                //支付时间
	PickUpTime  time.Time `json:"pick_up_time" gorm:";comment:取货时间;null"`                //取货时间
	ReturnTime  time.Time `json:"return_time" gorm:"comment:归还时间;null"`                  //归还时间
}
type OrderFrontList struct {
	MemberId int `form:"member_id"`
}

type AddOrder struct {
	MemberId      int     `json:"member_id" gorm:"type:bigint;comment:用户ID;not null"`
	OrderSn       string  `json:"order_sn" gorm:"type:varchar(191);comment:订单编号;not null"`
	TotalAmount   float64 `json:"total_amount" gorm:"type:decimal(10, 2) ;comment:订单总金额;not null"`
	FreightAmount float64 `json:"freight_amount" gorm:"type:decimal(10, 2) ;comment:运费金额;not null"` //运费金额
	ProductAmount float64 `json:"pay_amount" gorm:"type:decimal(10, 2) ;comment:商品金额;not null"`     //商品金额
	PledgeAmount  float64 `json:"pledge_amount" gorm:"type:decimal(10, 2) ;comment:押金;not null"`    //押金

	BookedTime    time.Time ` json:"booked_time" gorm:"comment:预约时间;null;"` //预约时间
	ReceiverName  string    `json:"receiver_name" gorm:"type:varchar(191);comment:收货人姓名;not null"`
	ReceiverPhone string    `json:"receiver_phone" gorm:"type:varchar(191);comment:收货人电话;not null"`
	OrderType     string    `json:"order_type" gorm:"type:varchar(191);comment:订单类型;not null"` //0->到店 1->送货上门
	Address       string    `json:"address" gorm:"type:varchar(191);comment:详细地址;not null"`

	PayType string `json:"pay_type" gorm:"type:varchar(191);comment:支付方式;not null"`
	Status  string `json:"status" gorm:"type:varchar(191);comment:订单状态;not null"` //0->待支付 1->已支付  2->待取货 3->租赁中 4->已归还

	PaymentTime time.Time `json:"payment_time" gorm:";comment:支付时间;null"` //支付时间

	Note string `json:"note" gorm:"type:varchar(191);comment:订单备注;not null"`
	Sku  []Sku  `json:"sku"`
}
type Sku struct {
	OrderId int `json:"order_id" `
	SkuId   int `json:"sku_id" `
	Count   int `json:"count" `
}
