package app

import (
	"fmt"

	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/containerhelper"
	"github.com/pandudpn/shopping-cart/app/container/servicecontainer"
	"github.com/pandudpn/shopping-cart/src/api/routes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func App() {
	port := fmt.Sprintf(":%d", viper.GetInt("application.port"))

	c, err := buildContainer()
	if err != nil {
		panic(err)
	}

	userController, err := containerhelper.GetUserController(c)
	if err != nil {
		panic(err)
	}

	cachedMiddleware, err := containerhelper.GetCachedMiddleware(c)
	if err != nil {
		panic(err)
	}

	routes := routes.RouteHandler{User: userController, Cached: cachedMiddleware}
	router := routes.Route()

	logrus.Fatal(router.Start(port))
}

func buildContainer() (container.Container, error) {
	factoryMap := make(map[string]interface{})
	c := servicecontainer.ServiceContainer{Factory: factoryMap}

	return &c, nil
}
