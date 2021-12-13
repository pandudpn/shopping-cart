package checkoutusecase

import (
	"context"

	"github.com/pandudpn/shopping-cart/src/api/presenter/checkoutpresenter"
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/helpers/processor"
	"github.com/pandudpn/shopping-cart/src/utils"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (cu *CheckoutUseCase) Update(ctx context.Context, req *model.RequestCheckout) utils.ResponseInterface {
	isCheckoutProgress := true

	cart, err := cu.getActiveCart(req.CartKey, req.UserId, isCheckoutProgress)
	if err != nil {
		logger.Log.Errorf("error get active cart %v", err)
		return checkoutpresenter.ResponseCheckout(isCheckoutProgress, nil, errActiveCart)
	}

	err = cu.TxRepo.TxEnd(func() error {
		err = cu.update(ctx, cart, req)
		if err != nil {
			return err
		}

		processor := processor.NewProcessor(cu.RedisRepo, cu.CartRepo, cu.CourierRepo, cu.UserAddressRepo, cu.PaymentMethodRepo)

		err = processor.Cart(cart, isCheckoutProgress)
		return err
	})

	if err != nil {
		return checkoutpresenter.ResponseCheckout(isCheckoutProgress, nil, errCheckout)
	}

	return checkoutpresenter.ResponseCheckout(isCheckoutProgress, cart, nil)
}

func (cu *CheckoutUseCase) update(ctx context.Context, cart *model.Cart, req *model.RequestCheckout) error {
	var err error
	if req.DeliveryAddress != nil || (req.DeliveryAddress != nil && req.DeliveryAddress.Id != 0) {
		userAddress, err := cu.UserAddressRepo.FindUserAddressById(req.DeliveryAddress.Id)
		if err != nil {
			return err
		}

		cart.UserAddressId = &userAddress.Id
		cart.SetUserAddress(userAddress)
	}

	if req.Courier != nil || (req.Courier != nil && req.Courier.Id != 0) {
		courier, err := cu.CourierRepo.FindCourierById(req.Courier.Id)
		if err != nil {
			return err
		}

		cart.CourierId = &courier.Id
		cart.SetCourier(courier)
	}

	if req.PaymentMethod != nil || (req.PaymentMethod != nil && req.PaymentMethod.Id != 0) {
		paymentMethod, err := cu.PaymentMethodRepo.FindPaymentMethodById(req.PaymentMethod.Id)
		if err != nil {
			return err
		}

		cart.PaymentMethodId = &paymentMethod.Id
		cart.SetPaymentMethod(paymentMethod)
	}

	err = cu.CartRepo.UpdateCart(cart)
	return err
}
