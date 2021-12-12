package paymenthandler

import (
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

type xenditVirtualAccount struct{}

func (xva *xenditVirtualAccount) Process(cart *model.Cart) (PaymentHandlerInterface, error) {
	return nil, errPaymentNotImplement
}
