package routes

import (
	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"
	"github.com/pandudpn/shopping-cart/src/api/controller"
	"github.com/pandudpn/shopping-cart/src/api/middleware"
)

type RouteHandler struct {
	User   controller.UserControllerInterface
	Cached middleware.CachedMiddlewareInterface
}

type RouteInterface interface {
	Route() *echo.Echo
}

func (rh *RouteHandler) Route() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Use(echoMiddleware.Logger())

	user := e.Group("/user")
	user.POST("", rh.User.LoginHandler)

	return e
}
