package model

import (
	"context"

	"gorm.io/gorm"
)

type Category struct {
	Base
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ProductList []Product `json:"product" gorm:"many2many:product_category"`
}

func (c Category) TableName() string {
	return "category"
}

type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

// preload ?
// c.db.WithContext 链路追踪
// .Model 指定查询的表
// .Where 条件
// .Preload 根据查出的category 以及关联表product_category,查出对应的ProductList
// .Find 返回结果绑定
func (c CategoryQuery) GetProductListByCategoryName(name string) (CategoryList []Category, err error) {
	err = c.db.WithContext(c.ctx).
		Model(&Category{}).
		Where(&Category{Name: name}).
		Preload("ProductList").
		Find(&CategoryList).Error
	return
}
