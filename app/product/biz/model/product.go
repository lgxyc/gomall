package model

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`
	// gorm:"many2many:product_category
	// 表示Product 和 Category 通过 product_category表实现多对多关联
	CategoryList []Category `json:"categoryList" gorm:"many2many:product_category"`
}

func (p Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

// p.db.WithContext 为了后续的链路追踪
func (p ProductQuery) GetById(productId int32) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, productId).Error
	return
}

func (p ProductQuery) SearchProcutList(q string) (productList []Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(productList, "name like ? or description like ?", "%"+q+"%", "%"+q+"%").Error
	return
}
