package service

import (
	"context"

	"github.com/lgxyc/gomall/app/product/biz/dal/mysql"
	"github.com/lgxyc/gomall/app/product/biz/model"
	product "github.com/lgxyc/gomall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)
	cList, err := categoryQuery.GetProductListByCategoryName(req.CategoryName)
	if err != nil {
		return nil, err
	}
	// db model -> rpc model
	var productList []*product.Product
	for _, vl := range cList {
		for _, v := range vl.ProductList {
			productList = append(productList, &product.Product{
				Id:          uint32(v.ID),
				Name:        v.Name,
				Description: v.Description,
				Picture:     v.Picture,
				Price:       v.Price,
			})
		}
	}
	return &product.ListProductsResp{ProductList: productList}, nil
}
