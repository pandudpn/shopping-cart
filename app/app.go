package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/containerhelper"
	"github.com/pandudpn/shopping-cart/app/container/servicecontainer"
	"github.com/pandudpn/shopping-cart/src/api/routes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func App() {
	loadConfig()

	port := fmt.Sprintf(":%d", viper.GetInt("application.port"))

	c, err := buildContainer()
	if err != nil {
		panic(err)
	}

	userController, err := containerhelper.GetUserController(c)
	if err != nil {
		panic(err)
	}

	routes := routes.RouteHandler{User: userController}
	router := routes.Route()

	logrus.Fatal(router.Start(port))
}

func buildContainer() (container.Container, error) {
	factoryMap := make(map[string]interface{})
	c := servicecontainer.ServiceContainer{Factory: factoryMap}

	return &c, nil
}

func loadConfig() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir = filepath.Join(dir, "../")
	configFile := fmt.Sprintf("%s/config.yml", dir)

	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)
	err = viper.ReadInConfig()
	if err != nil {
		logrus.Error(err)
	}
}
