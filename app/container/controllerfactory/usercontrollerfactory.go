package controllerfactory

import (
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/usecasefactory"
	"github.com/pandudpn/shopping-cart/src/api/controller/usercontroller"
	"github.com/pandudpn/shopping-cart/src/domain/usecase"
)

type userControllerFactory struct{}

func (ucf *userControllerFactory) Build(c container.Container) (ControllerFactoryInterface, error) {
	uucf, err := usecasefactory.GetUseCaseFbMap(constant.USER).Build(c)
	if err != nil {
		return nil, err
	}

	uuci := uucf.(usecase.UserUseCaseInterface)
	uc := usercontroller.UserController{UserUseCase: uuci}

	return &uc, nil
}
