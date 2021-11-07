package servicecontainer

import "github.com/pandudpn/shopping-cart/app/container/controllerfactory"

type ServiceContainer struct {
	Factory map[string]interface{}
}

func (sc *ServiceContainer) BuildController(code string) (interface{}, error) {
	return controllerfactory.GetControllerFbMap(code).Build(sc)
}

func (sc *ServiceContainer) Get(code string) (interface{}, bool) {
	value, found := sc.Factory[code]
	return value, found
}

func (sc *ServiceContainer) Put(code string, value interface{}) {
	sc.Factory[code] = value
}
