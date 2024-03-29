package cartusecase

import (
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/helpers/manager"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (cu *CartUseCase) getActiveCart(secretKey string, userId int, isCheckout bool) (*model.Cart, error) {
	cartManager := manager.NewCartManager(cu.CartRepo, cu.CartProductRepo, cu.ProductImageRepo, cu.UserRepo, cu.UserAddressRepo, cu.CourierRepo, cu.PaymentMethodRepo)

	cart, err := cartManager.GetActiveCart(secretKey, userId, isCheckout)
	if err != nil {
		return nil, err
	}

	logger.Log.Debug(cart.Id, cart.UserId)

	if cart.Id == 0 {
		err = cu.CartRepo.InsertNewCart(cart)
		if err != nil {
			return nil, err
		}
	}

	return cart, nil
}
