package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    uint32 `gorm:"type:int(11); not null; index:idx_user_id"`
	ProductId uint32 `gorm:"type:int(11);not null;"`
	Qty       uint32 `gorm:"type:int(11);not null;"`
}

func (c Cart) Tablename() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, cart *Cart) error {
	var row Cart
	//  先查找购物车中是否有该商品
	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{
			UserId:    cart.UserId,
			ProductId: cart.ProductId,
		}).
		First(&row).
		Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	// 如果存在更新数量
	if row.ID > 0 {
		return db.WithContext(ctx).
			Model(&Cart{}).
			Where(&Cart{
				UserId:    cart.UserId,
				ProductId: cart.ProductId,
			}).
			UpdateColumn("qty", row.Qty+cart.Qty).
			Error
	}
	return db.WithContext(ctx).Create(cart).Error
}

func ClearCart(ctx context.Context, db *gorm.DB, userId uint32) error {
	if userId == 0 {
		return errors.New("user id is required")
	}
	return db.WithContext(ctx).
		Delete(&Cart{}).
		Where("user_id = ?", userId).
		Error
}

func GetCartById(ctx context.Context, db *gorm.DB, userId uint32) ([]Cart, error) {
	var carts []Cart
	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where("user_id = ?", userId).
		Find(&carts).
		Error
	return carts, err
}
