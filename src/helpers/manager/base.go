package manager

import (
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

type CartManagerInterface interface {
	GetActiveCart(key string, userId int, isCheckout bool) (*model.Cart, error)
}

type StockManagerInterface interface {
}
