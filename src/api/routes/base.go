package routes

import (
	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"
	"github.com/pandudpn/shopping-cart/src/api/controller"
	"github.com/pandudpn/shopping-cart/src/api/middleware"
)

type RouteHandler struct {
	User    controller.UserControllerInterface
	Product controller.ProductControllerInterface
	Cart    controller.CartControllerInterface
	Cached  middleware.CachedMiddlewareInterface
}

type RouteInterface interface {
	Route() *echo.Echo
}

func (rh *RouteHandler) Route() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Use(echoMiddleware.Logger())

	auth := e.Group("/auth")
	auth.POST("/login", rh.User.LoginHandler)
	auth.POST("/register", rh.User.RegisterHandler)

	// product := e.Group("/product", echo.WrapMiddleware(rh.Cached.CachedData))
	product := e.Group("/product")
	product.GET("", rh.Product.GetProductsHandler)
	product.GET("/:id", rh.Product.DetailProductHandler)

	cart := e.Group("/cart", echo.WrapMiddleware(rh.Cached.CheckSession))
	cartProduct := cart.Group("/product")
	cartProduct.POST("/add", rh.Cart.AddToCartHandler)

	return e
}
