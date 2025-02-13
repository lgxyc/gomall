package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    int32 `gorm:"type:int(32); not null; index:idx_user_id"`
	ProductId int32 `gorm:"type:int(32);not null;"`
	Qty       int32 `gorm:"type:int(32);not null;"`
}

func (Cart) TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, c *Cart) error {
	var row Cart
	//  先查找购物车中是否有该商品
	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{
			UserId:    c.UserId,
			ProductId: c.ProductId,
		}).
		First(&row).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		return err
	}
	// 如果存在更新数量
	if row.ID > 0 {
		err = db.WithContext(ctx).
			Model(&Cart{}).
			Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).
			UpdateColumn("qty", gorm.Expr("qty+?", c.Qty)).Error
	} else {
		err = db.WithContext(ctx).Model(&Cart{}).Create(c).Error
	}
	return err
}

func ClearCart(ctx context.Context, db *gorm.DB, userId int32) error {
	if userId == 0 {
		return errors.New("user id is required")
	}
	return db.WithContext(ctx).
		Delete(&Cart{}, "user_id = ? ", userId).Error
}

func GetCartByUserId(ctx context.Context, db *gorm.DB, userId int32) (cartList []Cart, err error) {
	err = db.WithContext(ctx).
		Model(&Cart{}).
		Where("user_id = ?", userId).
		Find(&cartList).Error
	return cartList, err
}
