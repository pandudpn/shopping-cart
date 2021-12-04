package cartusecase

import (
	"errors"

	"github.com/pandudpn/shopping-cart/src/repository"
)

var (
	errProductNotFound = errors.New("cart.product.not_found")
	errQuantity        = errors.New("cart.product.quantity")
	errQueryInsert     = errors.New("query.insert.error")
	errActiveCart      = errors.New("cart.active.failed")
)

type CartUseCase struct {
	ProductRepo       repository.ProductRepositoryInterface
	CartRepo          repository.CartRepositoryInterface
	CartProductRepo   repository.CartProductRepositoryInterface
	ProductImageRepo  repository.ProductImageRepositoryInterface
	UserRepo          repository.UserRepositoryInterface
	UserAddressRepo   repository.UserAddressRepositoryInterface
	CourierRepo       repository.CourierRepositoryInterface
	PaymentMethodRepo repository.PaymentMethodRepositoryInterface
	TxRepo            repository.TxRepositoryInterface
}
