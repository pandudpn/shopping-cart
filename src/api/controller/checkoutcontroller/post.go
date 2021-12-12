package checkoutcontroller

import (
	"github.com/labstack/echo"
	"github.com/pandudpn/shopping-cart/src/api/middleware/cached"
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

func (cc *CheckoutController) PostHandler(e echo.Context) error {
	var (
		req     = e.Request()
		ctx     = req.Context()
		payload model.RequestCheckout
	)
	payload.CartKey = req.Header.Get("x-cart-key")
	payload.UserId = ctx.Value(cached.CtxUserId).(int)

	return cc.CheckoutUseCase.CreateOrder(ctx, &payload).JSON(e)
}
