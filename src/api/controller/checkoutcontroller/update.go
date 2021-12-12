package checkoutcontroller

import (
	"errors"

	"github.com/labstack/echo"
	"github.com/pandudpn/shopping-cart/src/api/middleware/cached"
	"github.com/pandudpn/shopping-cart/src/api/presenter/checkoutpresenter"
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (cc *CheckoutController) UpdateHandler(e echo.Context) error {
	var (
		req     = e.Request()
		ctx     = req.Context()
		payload model.RequestCheckout
	)
	payload.CartKey = req.Header.Get("x-cart-key")
	payload.UserId = ctx.Value(cached.CtxUserId).(int)

	err := e.Bind(&payload)
	if err != nil {
		logger.Log.Errorf("error bind payload %v", err)
		err = errors.New("body.payload")
		return checkoutpresenter.ResponseCheckout(true, nil, err).JSON(e)
	}

	return cc.CheckoutUseCase.Update(ctx, &payload).JSON(e)
}
