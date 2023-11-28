package model

import "time"

type Order struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id" mapstructure:"id"`
	CreatedAt time.Time `json:"created_at" mapstructure:"-"`

	MemberId      int     `json:"member_id" gorm:"type:bigint;comment:用户ID"`
	OrderSn       string  `json:"order_sn" gorm:"type:varchar(191);comment:订单编号"`
	TotalAmount   float64 `json:"total_amount" gorm:"type:decimal(10, 2) ;comment:订单总金额"`
	FreightAmount float64 `json:"freight_amount" gorm:"type:decimal(10, 2) ;comment:运费金额"` //运费金额
	ProductAmount float64 `json:"pay_amount" gorm:"type:decimal(10, 2) ;comment:商品金额"`     //商品金额
	PledgeAmount  float64 `json:"pledge_amount" gorm:"type:decimal(10, 2) ;comment:押金"`    //押金

	BookedTime    time.Time ` json:"booked_time" gorm:"comment:预约时间;default:null;"` //预约时间
	ReceiverName  string    `json:"receiver_name" gorm:"type:varchar(191);comment:收货人姓名"`
	ReceiverPhone string    `json:"receiver_phone" gorm:"type:varchar(191);comment:收货人电话"`
	OrderType     string    `json:"order_type" gorm:"type:varchar(191);comment:订单类型"` //0->到店 1->送货上门
	Address       string    `json:"address" gorm:"type:varchar(191);comment:详细地址"`

	PayType string `json:"pay_type" gorm:"type:varchar(191);comment:支付方式"`
	Status  string `json:"status" gorm:"type:varchar(191);comment:订单状态"` //0->待支付 1->已支付  2->待取货 3->租赁中 4->已归还

	PaymentTime time.Time `json:"payment_time" gorm:"comment:支付时间;default:null"` //支付时间
	PickUpTime  time.Time `json:"pick_up_time" gorm:"comment:取货时间;default:null"` //取货时间
	ReturnTime  time.Time `json:"return_time" gorm:"comment:归还时间;default:null"`  //归还时间

	Note string `json:"note" gorm:"type:varchar(191);comment:订单备注"`
}
