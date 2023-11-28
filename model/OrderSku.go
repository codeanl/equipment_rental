package model

type OrderSku struct {
	OrderId int `json:"order_id"`
	SkuId   int `json:"sku_id"`
	Count   int `json:"count" `
}
