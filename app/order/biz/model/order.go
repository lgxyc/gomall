package model

import (
	"context"

	"gorm.io/gorm"
)

// 收货人
type Consignee struct {
	Email         string
	StreetAddress string // 街道地址
	City          string // 城市
	State         string //	州
	Country       string // 国家
	ZipCode       string // 邮政编码
}
type Order struct {
	gorm.Model
	OrderId      string    `gorm:"type:varchar(100);uniqueIndex"`
	UserId       int32     `gorm:"type:int"`
	UserCurrency string    `gorm:"type:varchar(10)"`
	Consignee    Consignee `gorm:"embedded"` // embedded 嵌入结构体
	// gorm 一对多关联
	OrderItems []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
}

func (o Order) TableName() string {
	return "order"
}

func ListOrders(ctx context.Context, db *gorm.DB, userId int32) ([]Order, error) {
	var orders []Order
	// 关联表 preload
	err := db.WithContext(ctx).Where("user_id = ?", userId).Preload("OrderItems").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
