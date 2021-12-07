package checkoutcontroller

import (
	"github.com/labstack/echo"
	"github.com/pandudpn/shopping-cart/src/api/middleware/cached"
)

func (cc *CheckoutController) GetCheckoutHandler(e echo.Context) error {
	var (
		req = e.Request()
		ctx = req.Context()
		key = req.Header.Get("x-cart-key")
	)

	userId := ctx.Value(cached.CtxUserId).(int)

	return cc.CheckoutUseCase.GetCheckout(ctx, key, userId).JSON(e)
}
