package checkoutusecase

import (
	"context"

	"github.com/pandudpn/shopping-cart/src/api/presenter/checkoutpresenter"
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/helpers/manager"
	"github.com/pandudpn/shopping-cart/src/helpers/processor"
	"github.com/pandudpn/shopping-cart/src/utils"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (cu *CheckoutUseCase) GetCheckout(ctx context.Context, key string, userId int) utils.ResponseInterface {
	isCheckoutProgress := false
	isCheckout := true

	cart, err := cu.getActiveCart(key, userId, isCheckout)
	if err != nil {
		logger.Log.Errorf("error get active cart %v", err)
		return checkoutpresenter.ResponseCheckout(isCheckoutProgress, nil, errActiveCart)
	}

	err = cu.TxRepo.TxEnd(func() error {
		processor := processor.NewProcessor(cu.CartRepo, cu.CourierRepo, cu.UserAddressRepo, cu.PaymentMethodRepo)

		err = processor.Cart(cart, isCheckoutProgress)
		return err
	})

	if err != nil {
		logger.Log.Errorf("error process the cart %v", err)
		return checkoutpresenter.ResponseCheckout(isCheckoutProgress, nil, errCheckout)
	}

	return checkoutpresenter.ResponseCheckout(isCheckoutProgress, cart, nil)
}

func (cu *CheckoutUseCase) getActiveCart(key string, userId int, isCheckout bool) (*model.Cart, error) {
	cartManager := manager.NewCartManager(cu.CartRepo, cu.CartProductRepo, cu.ProductImageRepo, cu.UserRepo, cu.UserAddressRepo, cu.CourierRepo, cu.PaymentMethodRepo)

	cart, err := cartManager.GetActiveCart(key, userId, isCheckout)
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
