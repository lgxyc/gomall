package service

import (
	"context"

	"github.com/lgxyc/gomall/app/product/biz/dal/mysql"
	"github.com/lgxyc/gomall/app/product/biz/model"
	product "github.com/lgxyc/gomall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	productList, err := productQuery.SearchProcutList(req.Query)
	if err != nil {
		return nil, err
	}
	var productListResp []*product.Product
	for _, v := range productList {
		productListResp = append(productListResp, &product.Product{
			Id:          uint32(v.ID),
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
		})
	}

	return &product.SearchProductsResp{ProductList: productListResp}, nil
}
