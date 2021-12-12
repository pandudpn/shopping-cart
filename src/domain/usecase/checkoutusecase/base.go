package checkoutusecase

import (
	"errors"

	"github.com/pandudpn/shopping-cart/src/repository"
)

var (
	errCheckout      = errors.New("checkout.get.failed")
	errActiveCart    = errors.New("cart.active.failed")
	errCreatePayment = errors.New("payment.create.error")
	errInsert        = errors.New("query.insert.error")
	errUpdate        = errors.New("query.update.failed")
)

type CheckoutUseCase struct {
	CartRepo          repository.CartRepositoryInterface
	CartProductRepo   repository.CartProductRepositoryInterface
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
