package tables

import (
	"github.com/agilab/gostone/agiorm"
	"time"
)

type ShenzuanOrder struct {
	agiorm.BasicUintModel

	// 店铺 ID
	ShopId uint64 `gorm:"index:idx_shenzuan_order"`

	Amount uint64 //订单金额

	ProductName string // 商品名称

	DurationTime uint64 // 持续时间，有效期，以天为单位

	PayTime *time.Time // 支付时间

	PayType uint64 // 支付类型

	Status uint64 //1未支付，2支付成功已关闭，3 支付失败已关闭

	Deadline *time.Time //服务到期时间

	ShopName string

	TradeNo           string `json:"trade_no"`            // 支付宝交易号
	TotalAmount       uint64 `json:"total_amount"`        // 支付宝回调显示的订单金额 ，以分为单位
	ReceiptAmount     uint64 `json:"receipt_amount"`      // 支付宝回调显示的实收金额 ，以分为单位
}

func (ShenzuanOrder) TableName() string {
	return "shenzuan_order"
}

type Pay_Type int //支付类型，1支付宝，2，服务市场支付，3线下对公支付
const (
	Pay_Type_None           Pay_Type = iota
	Pay_Type_AliPay                  = 1
	Pay_Type_Service_Market          = 2
	Pay_Type_OffLine                 = 3
)

type Order_Ststus int //1未支付，2支付成功，订单已完成，3 订单已关闭
const (
	order_ststus_none         Order_Ststus = iota
	Order_Ststus_Not_Pay                   = 1
	Order_Ststus_Pay_Complete              = 2
	Order_Ststus_Closed                    = 3
)
