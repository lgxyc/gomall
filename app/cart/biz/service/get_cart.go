package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/lgxyc/gomall/app/cart/biz/dal/mysql"
	"github.com/lgxyc/gomall/app/cart/biz/model"
	cart "github.com/lgxyc/gomall/rpc_gen/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	cartList, err := model.GetCartById(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50002, err.Error())
	}
	var itemList []*cart.CartItem
	for _, item := range cartList {
		itemList = append(itemList, &cart.CartItem{
			ProductId: item.ProductId,
			Quantity:  item.Qty,
		})
	}
	return &cart.GetCartResp{ItemList: itemList}, nil
}
