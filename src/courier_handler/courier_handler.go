package courierhandler

import (
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

var handler = map[courier]courierHandlerFb{
	Shipper: &shipper{},
}

type courierHandlerFb interface {
	Process(order *model.Order) error
}

func GetCourierHandler(courierPartyKey courier) courierHandlerFb {
	return handler[courierPartyKey]
}
