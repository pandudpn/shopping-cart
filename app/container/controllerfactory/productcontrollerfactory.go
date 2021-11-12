package controllerfactory

import (
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/usecasefactory"
	"github.com/pandudpn/shopping-cart/src/api/controller/productcontroller"
	"github.com/pandudpn/shopping-cart/src/domain/usecase"
)

type productControllerFactory struct{}

func (pcf *productControllerFactory) Build(c container.Container) (ControllerFactoryInterface, error) {
	code := constant.PRODUCT

	puui, err := usecasefactory.GetUseCaseFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}

	puu := puui.(usecase.ProductUseCaseInterface)
	pc := productcontroller.ProductController{ProductUseCase: puu}

	return &pc, nil
}
