package calculator

import (
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/repository"
)

type cartCalcualtor struct {
	cartRepo repository.CartRepositoryInterface
}

func NewCartCalculator(cr repository.CartRepositoryInterface) CartCalculatorInterface {
	return &cartCalcualtor{
		cartRepo: cr,
	}
}

func (cc *cartCalcualtor) Calculate(cart *model.Cart) {
	total := cart.TotalProductsPrice + cart.TotalDeliveryCost

	cart.Total = total
}
