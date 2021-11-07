package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/pandudpn/shopping-cart/app"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	initConfig()

	app.App()
}

func initConfig() {
	_, b, _, _ := runtime.Caller(0)
	rootpath := filepath.Dir(b)
	configFile := fmt.Sprintf("%s/config.yml", rootpath)

	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error(err)
	}
}
