package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/lgxyc/gomall/app/checkout/infrac/rpc"
	rpccart "github.com/lgxyc/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/lgxyc/gomall/rpc_gen/kitex_gen/checkout"
	rpccheckout "github.com/lgxyc/gomall/rpc_gen/kitex_gen/checkout"
	rpcpayment "github.com/lgxyc/gomall/rpc_gen/kitex_gen/payment"
	rpcproduct "github.com/lgxyc/gomall/rpc_gen/kitex_gen/product"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	getCartResp, err := rpc.CartClient.GetCart(s.ctx,
		&rpccart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if getCartResp == nil || getCartResp.Cart == nil {
		return nil, kerrors.NewGRPCBizStatusError(5005002, "cart is empty")
	}
	var total float32

	for _, item := range getCartResp.Cart.Items {
		getProductResp, err := rpc.ProductClient.GetProduct(
			s.ctx, &rpcproduct.GetProductReq{Id: item.ProductId})
		if err != nil {
			return nil, err
		}
		if getProductResp.Product == nil {
			continue
		}
		total += getProductResp.Product.Price * float32(item.Quantity)
	}
	var orderId string
	u, _ := uuid.NewRandom()
	orderId = u.String()
	payReq := &rpcpayment.ChargeReq{
		OrderId: orderId,
		UserId:  req.UserId,
		Amount:  total,
		CreditCard: &rpcpayment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCartExpirationYear:  req.CreditCard.CreditCartExpirationYear,
			CreditCartExpirationMonth: req.CreditCard.CreditCartExpirationMonth,
		},
	}
	_, err = rpc.CartClient.ClearCart(
		s.ctx, &rpccart.ClearCartReq{
			UserId: req.UserId,
		})
	if err != nil {
		klog.Error(err.Error())
	}
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}
	klog.Info(paymentResult)
	resp = &rpccheckout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}

	return
}
