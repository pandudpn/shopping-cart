package controllerfactory

import (
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/usecasefactory"
	"github.com/pandudpn/shopping-cart/src/api/controller/checkoutcontroller"
	"github.com/pandudpn/shopping-cart/src/domain/usecase"
)

type checkoutControllerFactory struct{}

func (ccf *checkoutControllerFactory) Build(c container.Container) (ControllerFactoryInterface, error) {
	code := constant.CHECKOUT

	cusi, err := usecasefactory.GetUseCaseFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}

	cuui := cusi.(usecase.CheckoutUseCaseInterface)
	cc := checkoutcontroller.CheckoutController{CheckoutUseCase: cuui}

	return &cc, nil
}
