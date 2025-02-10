package model

import "gorm.io/gorm"

// 订单 商品关联表
type OrderItem struct {
	gorm.Model
	OrderIdRefer string  `gorm:"type:varchar(100);index"`
	ProductId    int32   `gorm:"type:int(11)"`
	Quantity     int32   `gorm:"type:int(11)"`
	Cost         float32 `gorm:"type:decimal(10,2)"`
}

func (o OrderItem) TableName() string {
	return "order_item"
}
