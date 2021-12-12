package checkoutusecase

import (
	"context"

	"github.com/pandudpn/shopping-cart/src/api/presenter/checkoutpresenter"
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/helpers/processor"
	"github.com/pandudpn/shopping-cart/src/utils"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (cu *CheckoutUseCase) CreateOrder(ctx context.Context, req *model.RequestCheckout) utils.ResponseInterface {
	var order *model.Order
	isCheckoutProgress := true

	cart, err := cu.getActiveCart(req.CartKey, req.UserId)
	if err != nil {
		logger.Log.Errorf("error get active cart %v", err)
		return checkoutpresenter.FinishCheckout(nil, errActiveCart)
	}

	err = cu.TxRepo.TxEnd(func() error {
		processor := processor.NewProcessor(cu.CartRepo, cu.CourierRepo, cu.UserAddressRepo, cu.PaymentMethodRepo)

		err = processor.Cart(cart, isCheckoutProgress)
		if err != nil {
			return err
		}

		order, err = cu.convertCartToOrder(cart)
		return err
	})

	return checkoutpresenter.FinishCheckout(order, err)
}
