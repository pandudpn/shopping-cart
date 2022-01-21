package callbackusecase

import (
	"errors"

	"github.com/pandudpn/shopping-cart/src/repository"
)

var (
	errOrderNotFound = errors.New("order.notfound")
	errGetRelation   = errors.New("database.relation.error")
	errPaymentFailed = errors.New("payment.failed")
	errUpdateOrder   = errors.New("order.update.error")
)

type CallbackUseCase struct {
	// user, and courier related
	ProductRepo       repository.ProductRepositoryInterface
	ProductImageRepo  repository.ProductImageRepositoryInterface
	UserRepo          repository.UserRepositoryInterface
	UserAddressRepo   repository.UserAddressRepositoryInterface
	CourierRepo       repository.CourierRepositoryInterface
	PaymentMethodRepo repository.PaymentMethodRepositoryInterface
	TxRepo            repository.TxRepositoryInterface
	// Order related
	OrderRepo               repository.OrderRepositoryInterface
	OrderProductRepo        repository.OrderProductRepositoryInterface
	OrderPaymentRepo        repository.OrderPaymentRepositoryInterface
	OrderDeliveryRepo       repository.OrderDeliveryRepositoryInterface
	OrderDeliveryStatusRepo repository.OrderDeliveryStatusRepositoryInterface
}
