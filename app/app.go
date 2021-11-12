package app

import (
	"fmt"

	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/containerhelper"
	"github.com/pandudpn/shopping-cart/app/container/logfactory"
	"github.com/pandudpn/shopping-cart/app/container/servicecontainer"
	"github.com/pandudpn/shopping-cart/src/api/routes"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
	"github.com/spf13/viper"
)

func App() {
	initLogger()

	port := fmt.Sprintf(":%d", viper.GetInt("application.port"))

	c, err := buildContainer()
	if err != nil {
		panic(err)
	}

	userController, err := containerhelper.GetUserController(c)
	if err != nil {
		logger.Log.Error(err)
		panic(err)
	}

	productController, err := containerhelper.GetProductController(c)
	if err != nil {
		logger.Log.Error(err)
		panic(err)
	}

	cachedMiddleware, err := containerhelper.GetCachedMiddleware(c)
	if err != nil {
		logger.Log.Error(err)
		panic(err)
	}

	routes := routes.RouteHandler{User: userController, Cached: cachedMiddleware, Product: productController}
	router := routes.Route()

	logger.Log.Fatal(router.Start(port))
}

func buildContainer() (container.Container, error) {
	factoryMap := make(map[string]interface{})
	c := servicecontainer.ServiceContainer{Factory: factoryMap}

	return &c, nil
}

func initLogger() {
	logfactory.GetLogFb().Build()
}
