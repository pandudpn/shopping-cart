package controllerfactory

import (
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/usecasefactory"
	"github.com/pandudpn/shopping-cart/src/api/controller/callbackcontroller"
	"github.com/pandudpn/shopping-cart/src/domain/usecase"
)

type callbackControllerFactory struct{}

func (ccf *callbackControllerFactory) Build(c container.Container) (ControllerFactoryInterface, error) {
	code := constant.CALLBACK

	cusi, err := usecasefactory.GetUseCaseFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}

	cuui := cusi.(usecase.CallbackUseCaseInterface)
	cc := callbackcontroller.CallbackController{
		CallbackUseCase: cuui,
	}

	return &cc, nil
}
