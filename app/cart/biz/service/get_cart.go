package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/lgxyc/gomall/app/cart/biz/dal/mysql"
	"github.com/lgxyc/gomall/app/cart/biz/model"
	rpccart "github.com/lgxyc/gomall/rpc_gen/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *rpccart.GetCartReq) (resp *rpccart.GetCartResp, err error) {
	cartList, err := model.GetCartByUserId(s.ctx, mysql.DB, req.GetUserId())
	if err != nil {
		return nil, kerrors.NewBizStatusError(50002, err.Error())
	}
	var items []*rpccart.CartItem
	for _, v := range cartList {
		items = append(items, &rpccart.CartItem{ProductId: v.ProductId, Quantity: int32(v.Qty)})
	}

	return &rpccart.GetCartResp{Cart: &rpccart.Cart{UserId: req.GetUserId(), Items: items}}, nil
}
