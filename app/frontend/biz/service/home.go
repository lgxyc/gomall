package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/lgxyc/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/lgxyc/gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/lgxyc/gomall/rpc_gen/kitex_gen/product"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	p, err := rpc.ProductClient.ListProducts(h.Context, &rpcproduct.ListProductsReq{})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"Title": "Hot sales",
		"Items": p.ProductList,
	}, err
}
