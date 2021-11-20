package controllerfactory

import (
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/usecasefactory"
	"github.com/pandudpn/shopping-cart/src/api/controller/cartcontroller"
	"github.com/pandudpn/shopping-cart/src/domain/usecase"
)

type cartControllerFactory struct{}

func (ccf *cartControllerFactory) Build(c container.Container) (ControllerFactoryInterface, error) {
	code := constant.CART

	cusi, err := usecasefactory.GetUseCaseFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}

	cuui := cusi.(usecase.CartUseCaseInterface)
	cc := cartcontroller.CartController{CartUsecase: cuui}

	return &cc, nil
}
