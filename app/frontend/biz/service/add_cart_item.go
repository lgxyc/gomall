package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	cart "github.com/lgxyc/gomall/app/frontend/hertz_gen/frontend/cart"
	common "github.com/lgxyc/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/lgxyc/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/lgxyc/gomall/app/frontend/utils"
	rpccart "github.com/lgxyc/gomall/rpc_gen/kitex_gen/cart"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp *common.Empty, err error) {
	_, err = rpc.CartClient.AddItem(h.Context,
		&rpccart.AddItemReq{
			UserId: frontendutils.GetUserIdFromCtx(h.Context),
			Item: &rpccart.CartItem{
				ProductId: req.ProductId,
				Quantity:  int32(req.Quantity),
			},
		})
	if err != nil {
		return nil, err
	}
	return
}
