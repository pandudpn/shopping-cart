package middlewarefactory

import (
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
)

var mdFbMap = map[string]middlewareFbInterface{
	constant.REDIS: &redisMiddlewareFactory{},
}

type MiddlewareFactoryInterface interface{}

type middlewareFbInterface interface {
	Build(c container.Container) (MiddlewareFactoryInterface, error)
}

func GetMiddlewareFbMap(code string) middlewareFbInterface {
	return mdFbMap[code]
}
