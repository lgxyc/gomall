package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/lgxyc/gomall/app/cart/biz/dal/mysql"
	"github.com/lgxyc/gomall/app/cart/biz/model"
	cart "github.com/lgxyc/gomall/rpc_gen/kitex_gen/cart"
)

type ClearCartService struct {
	ctx context.Context
} // NewClearCartService new ClearCartService
func NewClearCartService(ctx context.Context) *ClearCartService {
	return &ClearCartService{ctx: ctx}
}

// Run create note info
func (s *ClearCartService) Run(req *cart.ClearCartReq) (resp *cart.ClearCartResp, err error) {
	err = model.ClearCart(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50001, err.Error())
	}
	return &cart.ClearCartResp{}, nil
}
