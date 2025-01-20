package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/lgxyc/gomall/app/cart/biz/dal/mysql"
	"github.com/lgxyc/gomall/app/cart/biz/model"
	"github.com/lgxyc/gomall/app/cart/infra/rpc"
	cart "github.com/lgxyc/gomall/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/lgxyc/gomall/rpc_gen/kitex_gen/product"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	p, err := rpc.ProductClient.GetProduct(s.ctx, &rpcproduct.GetProductReq{
		Id: req.Item.ProductId,
	})
	if err != nil {
		return nil, err
	}
	if p.Product == nil || p.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product not found")
	}
	newcart := &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       req.Item.Quantity,
	}
	err = model.AddItem(s.ctx, mysql.DB, newcart)
	if err != nil {
		return nil, kerrors.NewBizStatusError(5000, err.Error())
	}

	return
}
