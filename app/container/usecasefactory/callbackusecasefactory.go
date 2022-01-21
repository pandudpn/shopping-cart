package usecasefactory

import (
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/repositoryfactory"
	"github.com/pandudpn/shopping-cart/src/domain/usecase/callbackusecase"
	"github.com/pandudpn/shopping-cart/src/repository"
)

type callbackUseCaseFactory struct{}

func (cuuf *callbackUseCaseFactory) Build(c container.Container) (UseCaseFactoryInterface, error) {
	// PRODUCT REPOSITORY
	code := constant.PRODUCT
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

	// TX Repository
	code = constant.TX
	rufiTx, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	uriTx := rufiTx.(repository.TxRepositoryInterface)

	// Order Repository
	code = constant.ORDER
	orfi, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	ori := orfi.(repository.OrderRepositoryInterface)

	// Order Product Repository
	code = constant.ORDERPRODUCT
	orpfi, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	orpi := orpfi.(repository.OrderProductRepositoryInterface)

	// Order Payment Repository
	code = constant.ORDERPAYMENT
	orppfi, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	orppi := orppfi.(repository.OrderPaymentRepositoryInterface)

	// Order Delivery Repository
	code = constant.ORDERDELIVERY
	odrfi, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	odri := odrfi.(repository.OrderDeliveryRepositoryInterface)

	// Order Delivery Status Repository
	code = constant.ORDERDELIVERYSTATUS
	odsrfi, err := repositoryfactory.GetRepositoryFbMap(code).Build(c)
	if err != nil {
		return nil, err
	}
	odsri := odsrfi.(repository.OrderDeliveryStatusRepositoryInterface)

	cuc := callbackusecase.CallbackUseCase{
		ProductRepo:             pri,
		ProductImageRepo:        pir,
		UserRepo:                uri,
		UserAddressRepo:         uari,
		CourierRepo:             cori,
		PaymentMethodRepo:       pmri,
		TxRepo:                  uriTx,
		OrderRepo:               ori,
		OrderProductRepo:        orpi,
		OrderPaymentRepo:        orppi,
		OrderDeliveryRepo:       odri,
		OrderDeliveryStatusRepo: odsri,
	}

	return &cuc, nil
}
