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
	// CART REPOSITORY
	code := constant.CART
	crif, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	cri := crif.(repository.CartRepositoryInterface)

	// CART PRODUCT REPOSITORY
	code = constant.CART_PRODUCT
	cprif, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	cpri := cprif.(repository.CartProductRepositoryInterface)

	// PRODUCT REPOSITORY
	code = constant.PRODUCT
	prif, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	pri := prif.(repository.ProductRepositoryInterface)

	// Product Image Repository
	code = constant.PRODUCTIMAGE
	pirif, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	pir := pirif.(repository.ProductImageRepositoryInterface)

	// User Repository
	code = constant.USER
	urif, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	uri := urif.(repository.UserRepositoryInterface)

	// User Address Repository
	code = constant.USERADDRESS
	uarif, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	uari := uarif.(repository.UserAddressRepositoryInterface)

	// Courier Repository
	code = constant.COURIER
	corif, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	cori := corif.(repository.CourierRepositoryInterface)

	// Payment Method Repository
	code = constant.PAYMENTMETHOD
	pmrif, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	pmri := pmrif.(repository.PaymentMethodRepositoryInterface)

	code = constant.TX

	rufiTx, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}

	uriTx := rufiTx.(repository.TxRepositoryInterface)

	cuc := cartusecase.CartUseCase{
		ProductRepo:       pri,
		CartRepo:          cri,
		CartProductRepo:   cpri,
		ProductImageRepo:  pir,
		UserRepo:          uri,
		UserAddressRepo:   uari,
		CourierRepo:       cori,
		PaymentMethodRepo: pmri,
		TxRepo:            uriTx,
	}

	return &cuc, nil
}
