package containerhelper

import (
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/src/api/controller"
	"github.com/pandudpn/shopping-cart/src/api/middleware"
)

func GetUserController(c container.Container) (controller.UserControllerInterface, error) {
	key := "user"

	uci, err := c.BuildController(key)
	if err != nil {
		return nil, err
	}

	return uci.(controller.UserControllerInterface), nil
}

func GetCachedMiddleware(c container.Container) (middleware.CachedMiddlewareInterface, error) {
	mfi, err := c.BuildMiddleware(constant.CACHED)
	if err != nil {
		return nil, err
	}

	return mfi.(middleware.CachedMiddlewareInterface), nil
}
