package paymenthandler

import "errors"

const (
	OneTimePayment = "ONE_TIME_PAYMENT"
	Product        = "PRODUCT"
	IDR            = "IDR"
	// E-wallet
	OVO     = "id_ovo"
	SHOPEE  = "id_shopeepay"
	LINKAJA = "id_linkaja"
	DANA    = "id_dana"
)

var (
	errPayment             = errors.New("payment.create.error")
	errPaymentNotImplement = errors.New("payment_method.not_implemented")
)
