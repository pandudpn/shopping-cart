package processor

import (
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/helpers/calculator"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (p *processor) Cart(cart *model.Cart, isCheckoutOnProgress bool) error {
	var err error
	if !isCheckoutOnProgress {
		cart.CourierId = nil
		cart.PaymentMethodId = nil

		err = p.cartRepo.UpdateCart(cart)
		if err != nil {
			logger.Log.Errorf("error update cart %v", err)
			return ErrCartUpdate
		}
	}

	if cart.UserAddressId == nil {
		userAddress, err := p.userAddressRepo.FindDefaultDeliveryByUser(cart.GetUser())
		if err != nil {
			logger.Log.Errorf("error get default delivery address %v", err)
			return ErrDeliveryAddress
		}

		cart.SetUserAddress(userAddress)
		cart.UserAddressId = &userAddress.Id

		err = p.cartRepo.UpdateCart(cart)
		if err != nil {
			logger.Log.Errorf("error update cart %v", err)
			return ErrCartUpdate
		}
	}

	err = p.GetAvailableCourier(cart)
	if err != nil {
		return err
	}

	calculator := calculator.NewCartCalculator(p.cartRepo)
	calculator.Calculate(cart)

	return nil
}
