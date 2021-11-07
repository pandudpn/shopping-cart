package controllerfactory

import (
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
)

var cFbMap = map[string]controllerFbInterface{
	constant.USER: &userControllerFactory{},
}

type ControllerFactoryInterface interface{}

type controllerFbInterface interface {
	Build(c container.Container) (ControllerFactoryInterface, error)
}

func GetControllerFbMap(code string) controllerFbInterface {
	return cFbMap[code]
}
