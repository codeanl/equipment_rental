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

	Status      string    `json:"status" gorm:"type:varchar(191);comment:订单状态;not null"` //0->待支付 1->已支付  2->待取货 3->租赁中 4->已归还
	PaymentTime time.Time `json:"payment_time" gorm:";comment:支付时间;null"`                //支付时间
	PickUpTime  time.Time `json:"pick_up_time" gorm:";comment:取货时间;null"`                //取货时间
	ReturnTime  time.Time `json:"return_time" gorm:"comment:归还时间;null"`                  //归还时间
}
