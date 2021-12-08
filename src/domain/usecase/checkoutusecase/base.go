package checkoutusecase

import (
	"errors"

	"github.com/pandudpn/shopping-cart/src/repository"
)

var (
	errCheckout   = errors.New("checkout.get.failed")
	errActiveCart = errors.New("cart.active.failed")
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
}