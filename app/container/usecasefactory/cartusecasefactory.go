package usecasefactory

import (
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/repositoryfactory"
	"github.com/pandudpn/shopping-cart/src/domain/usecase/cartusecase"
	"github.com/pandudpn/shopping-cart/src/repository"
)

type cartUseCaseFactory struct{}

func (cucf *cartUseCaseFactory) Build(c container.Container) (UseCaseFactoryInterface, error) {
	code := constant.CART

	crif, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	cri := crif.(repository.CartRepositoryInterface)

	code = constant.CART_PRODUCT

	cprif, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	cpri := cprif.(repository.CartProductRepositoryInterface)

	code = constant.PRODUCT

	prif, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	pri := prif.(repository.ProductRepositoryInterface)

	code = constant.TX

	rufiTx, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}

	uriTx := rufiTx.(repository.TxRepositoryInterface)

	cuc := cartusecase.CartUseCase{
		ProductRepo:     pri,
		CartRepo:        cri,
		CartProductRepo: cpri,
		TxRepo:          uriTx,
	}

	return &cuc, nil
}
