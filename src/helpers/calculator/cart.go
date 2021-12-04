package calculator

import (
	"context"

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

func (cc *cartCalcualtor) Calculate(ctx context.Context, cart *model.Cart) error {
	return nil
}
