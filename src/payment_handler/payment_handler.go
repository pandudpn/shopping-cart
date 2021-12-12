package paymenthandler

import (
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

var handler = map[string]paymentHandlerFb{
	constant.XENDITEWALLET: &xenditEwallet{},
}

type PaymentHandlerInterface interface{}

type paymentHandlerFb interface {
	Process(cart *model.Cart) (PaymentHandlerInterface, error)
}

func GetHandlerPayment(paymentMethodKey string) paymentHandlerFb {
	return handler[paymentMethodKey]
}
