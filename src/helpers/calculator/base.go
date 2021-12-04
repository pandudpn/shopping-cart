package calculator

import (
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

type CartCalculatorInterface interface {
	Calculate(cart *model.Cart)
}
