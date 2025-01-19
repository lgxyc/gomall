package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	product "github.com/lgxyc/gomall/app/frontend/hertz_gen/frontend/product"
	"github.com/lgxyc/gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/lgxyc/gomall/rpc_gen/kitex_gen/product"
)

type SearchProductListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductListService(Context context.Context, RequestContext *app.RequestContext) *SearchProductListService {
	return &SearchProductListService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductListService) Run(req *product.SearchProductListReq) (resp map[string]any, err error) {
	p, err := rpc.ProductClient.SearchProducts(h.Context,
		&rpcproduct.SearchProductsReq{
			Query: req.Query,
		})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"Items": p.ProductList,
		"query": req.Query,
	}, nil
}
