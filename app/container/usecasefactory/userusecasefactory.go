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
	rufi, err := repositoryfactory.GetRepositoryFbMap(constant.USER).Build(c)
	if err != nil {
		return nil, err
	}

	uri := rufi.(repository.UserRepositoryInterface)

	rufiTx, err := repositoryfactory.GetRepositoryFbMap(constant.TX).Build(c)
	if err != nil {
		return nil, err
	}

	uriTx := rufiTx.(repository.TxRepositoryInterface)

	rrfi, err := repositoryfactory.GetRepositoryFbMap(constant.REDIS).Build(c)
	if err != nil {
		return nil, err
	}
	rri := rrfi.(repository.RedisRepositoryInterface)

	uuc := userusecase.UserUseCase{UserRepo: uri, TxRepo: uriTx, RedisRepo: rri}

	return &uuc, nil
}
