package service

import (
	"context"
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

	return
}
