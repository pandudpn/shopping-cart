package cartcontroller

import (
	"errors"

	"github.com/labstack/echo"
	"github.com/pandudpn/shopping-cart/src/api/middleware/cached"
	"github.com/pandudpn/shopping-cart/src/api/presenter/cartpresenter"
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (cc *CartController) AddToCartHandler(e echo.Context) error {
	var (
		req         = e.Request()
		ctx         = req.Context()
		payload     = &model.RequestAddToCart{}
		isAddToCart = true
	)

	if err := e.Bind(&payload); err != nil {
		logger.Log.Errorf("error parsing payload addtocart %v", err)
		err = errors.New("body.payload")
		return cartpresenter.ResponseCart(isAddToCart, nil, err).JSON(e)
	}

	payload.UserId = ctx.Value(cached.CtxUserId).(int)
	logger.Log.Debugf("payload add to cart %v", payload)

	return cc.CartUsecase.AddToCart(ctx, payload).JSON(e)
}
