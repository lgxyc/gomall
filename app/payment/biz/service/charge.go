package service

import (
	"context"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"github.com/lgxyc/gomall/app/payment/biz/dal/mysql"
	"github.com/lgxyc/gomall/app/payment/biz/model"
	payment "github.com/lgxyc/gomall/rpc_gen/kitex_gen/payment"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// 第三方库校验卡信息

	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCartExpirationYear)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCartExpirationYear)),
	}
	// 参数 true 表示可以使用特殊卡号
	err = card.Validate(true)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4004001, err.Error())
	}
	// 生成交易id
	transactionId, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005001, err.Error())
	}

	// 创建支付记录
	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserId:        int32(req.UserId),
		OrdrId:        req.OrderId,
		Amount:        req.Amount,
		TransactionId: transactionId.String(),
		PayAt:         time.Now(),
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005002, err.Error())
	}
	return &payment.ChargeResp{TransactionId: transactionId.String()}, nil
}
