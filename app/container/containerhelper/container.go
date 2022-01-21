package containerhelper

import (
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/src/api/controller"
	"github.com/pandudpn/shopping-cart/src/api/middleware"
)

func GetUserController(c container.Container) (controller.UserControllerInterface, error) {
	key := constant.USER

	uci, err := c.BuildController(key)
	if err != nil {
		return nil, err
	}

	return uci.(controller.UserControllerInterface), nil
}

func GetProductController(c container.Container) (controller.ProductControllerInterface, error) {
	key := constant.PRODUCT

	pci, err := c.BuildController(key)
	if err != nil {
		return nil, err
	}

	return pci.(controller.ProductControllerInterface), nil
}

func GetCartController(c container.Container) (controller.CartControllerInterface, error) {
	key := constant.CART

	cci, err := c.BuildController(key)
	if err != nil {
		return nil, err
	}

	return cci.(controller.CartControllerInterface), nil
}

func GetCheckoutController(c container.Container) (controller.CheckoutControllerInterface, error) {
	key := constant.CHECKOUT

	cci, err := c.BuildController(key)
	if err != nil {
		return nil, err
	}

	return cci.(controller.CheckoutControllerInterface), nil
}

func GetCallbackController(c container.Container) (controller.CallbackControllerInterface, error) {
	key := constant.CALLBACK

	cci, err := c.BuildController(key)
	if err != nil {
		return nil, err
	}

	return cci.(controller.CallbackControllerInterface), nil
}

func GetCachedMiddleware(c container.Container) (middleware.CachedMiddlewareInterface, error) {
	mfi, err := c.BuildMiddleware(constant.REDIS)
	if err != nil {
		return nil, err
	}

	return mfi.(middleware.CachedMiddlewareInterface), nil
}
