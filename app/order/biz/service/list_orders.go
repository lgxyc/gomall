package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/lgxyc/gomall/app/order/biz/dal/mysql"
	"github.com/lgxyc/gomall/app/order/biz/model"
	order "github.com/lgxyc/gomall/rpc_gen/kitex_gen/order"
)

type ListOrdersService struct {
	ctx context.Context
} // NewListOrdersService new ListOrdersService
func NewListOrdersService(ctx context.Context) *ListOrdersService {
	return &ListOrdersService{ctx: ctx}
}

// Run create note info
func (s *ListOrdersService) Run(req *order.ListOrdersReq) (resp *order.ListOrdersResp, err error) {
	// Finish your business logic.
	orders, err := model.ListOrders(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(500001, err.Error())
	}
	var OrderList []*order.Order
	for _, v := range orders {
		OrderList = append(OrderList, &order.Order{
			OrderId:      v.OrderId,
			UserId:       v.UserId,
			UserCurrency: v.UserCurrency,
			Address: &order.Address{
				StreetAddress: v.Consignee.StreetAddress,
				City:          v.Consignee.City,
				State:         v.Consignee.State,
				Country:       v.Consignee.Country,
				ZipCode:       v.Consignee.ZipCode,
			},
		})
	}
	resp = &order.ListOrdersResp{
		OrderList: OrderList,
	}

	return resp, nil
}
