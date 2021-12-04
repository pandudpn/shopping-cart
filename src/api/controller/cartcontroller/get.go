package cartcontroller

import (
	"github.com/labstack/echo"
	"github.com/pandudpn/shopping-cart/src/api/middleware/cached"
)

func (cc *CartController) GetCartHandler(e echo.Context) error {
	var (
		req = e.Request()
		ctx = req.Context()
		key = e.QueryParam("x-cart-key")
	)

	userId := ctx.Value(cached.CtxUserId).(int)

	return cc.CartUsecase.GetCart(ctx, userId, key).JSON(e)
}
