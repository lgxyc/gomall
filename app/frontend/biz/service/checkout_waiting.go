package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	checkout "github.com/lgxyc/gomall/app/frontend/hertz_gen/frontend/checkout"
	"github.com/lgxyc/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/lgxyc/gomall/app/frontend/utils"
	rpccheckout "github.com/lgxyc/gomall/rpc_gen/kitex_gen/checkout"
	rpcpayment "github.com/lgxyc/gomall/rpc_gen/kitex_gen/payment"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	UserId := frontendutils.GetUserIdFromCtx(h.Context)
	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    UserId,
		Email:     req.Email,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Address: &rpccheckout.Address{
			Country:       req.Country,
			ZipCode:       req.Zipcode,
			City:          req.City,
			StreetAddress: req.Street,
			State:         req.Province,
		},
		CreditCard: &rpcpayment.CreditCardInfo{
			CreditCardNumber:          req.CardNum,
			CreditCartExpirationYear:  req.ExpirationYear,
			CreditCartExpirationMonth: req.ExpirationMonth,
			CreditCardCvv:             req.Cvv,
		},
	})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"title":    "waiting",
		"redirect": "/checkout/result",
	}, nil
}
