package usecasefactory

import (
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/repositoryfactory"
	"github.com/pandudpn/shopping-cart/src/domain/usecase/userusecase"
	"github.com/pandudpn/shopping-cart/src/repository"
)

type userUseCaseFactory struct{}

func (uucf *userUseCaseFactory) Build(c container.Container) (UseCaseFactoryInterface, error) {
	rufi, err := repositoryfactory.GetRepositoryFbMap(constant.USER).Build(c, constant.ENABLETX)
	if err != nil {
		return nil, err
	}

	uri := rufi.(repository.UserRepositoryInterface)
	uuc := userusecase.UserUseCase{UserRepo: uri}

	return &uuc, nil
}
