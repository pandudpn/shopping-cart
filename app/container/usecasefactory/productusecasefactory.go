package usecasefactory

import (
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/repositoryfactory"
	"github.com/pandudpn/shopping-cart/src/domain/usecase/productusecase"
	"github.com/pandudpn/shopping-cart/src/repository"
)

type productUseCaseFactory struct{}

func (pucf *productUseCaseFactory) Build(c container.Container) (UseCaseFactoryInterface, error) {
	code := constant.PRODUCT
	prfi, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	pri := prfi.(repository.ProductRepositoryInterface)

	code = constant.PRODUCTCATEGORY

	pcrfi, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	pcri := pcrfi.(repository.ProductCategoryRepositoryInterface)

	puc := productusecase.ProductUseCase{ProductRepo: pri, CategoryRepo: pcri}

	return &puc, nil
}
