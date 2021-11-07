package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pandudpn/shopping-cart/src/api/controller"
)

type RouteHandler struct {
	User controller.UserControllerInterface
}

type RouteInterface interface {
	Route() *echo.Echo
}

func (rh *RouteHandler) Route() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())

	user := e.Group("/user")
	user.POST("", rh.User.LoginHandler)

	return e
}
