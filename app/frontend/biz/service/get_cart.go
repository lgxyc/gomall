package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/lgxyc/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/lgxyc/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/lgxyc/gomall/app/frontend/utils"
	rpccart "github.com/lgxyc/gomall/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/lgxyc/gomall/rpc_gen/kitex_gen/product"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	rpcResp, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: frontendutils.GetUserIdFormCtx(h.Context),
	})
	if err != nil {
		return nil, err
	}
	var productList []map[string]string
	var total float64
	for _, item := range rpcResp.Cart.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
			Id: item.ProductId,
		})
		if err != nil {
			continue
		}
		p := productResp.Product
		productList = append(productList, map[string]string{
			"Name":        p.Name,
			"Descritpion": p.Description,
			"Price":       strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture":     p.Picture,
			"Qty":         strconv.Itoa(int(item.Quantity)),
		})
		total += float64(p.Price) * float64(item.Quantity)
	}

	return utils.H{
		"Title":       "Cart",
		"ProductList": productList,
		"Total":       total,
	}, nil
}
