package processor

import (
	"crypto/tls"
	"errors"
	"net/http"

	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/repository"
)

// Kumpulan-kumpulan error yang nanti nya akan di convert pada presenter
var (
	ErrCourierNotAvail = errors.New("courier.not_avail")
	ErrCourier         = errors.New("courier.error")
	ErrCartUpdate      = errors.New("cart.update.error")
	ErrDeliveryAddress = errors.New("checkout.delivery_address.not_found")
	ErrPaymentMethod   = errors.New("checkout.payment_method.error")
)

type processor struct {
	client            *http.Client
	cartRepo          repository.CartRepositoryInterface
	courierRepo       repository.CourierRepositoryInterface
	userAddressRepo   repository.UserAddressRepositoryInterface
	paymentMethodRepo repository.PaymentMethodRepositoryInterface
	redisRepo         repository.RedisRepositoryInterface
}

// ProcessorInterface adalah sebuah interface yang menampung method yang dapat di akses pada package 'Processor' ini
type ProcessorInterface interface {
	// GetAvailableCourier digunakan untuk mengambil list-list courier yg tersedia ketika melakukan checkout
	// pada method ini juga bisa digunakan untuk mengambil data courier yang sudah dipilih
	// lalu akan di convert pada presenter nantinya
	GetAvailableCourier(cart *model.Cart) error
	// Cart method untuk memproses seluruh data dari availableCourier, availablePaymentMethod, dan juga calculator
	// dari cart itu sendiri. method ini akan digunakan ketika user melakukan checkout
	Cart(cart *model.Cart, isCheckoutOnProgress bool) error
	// GetAvailablePaymentMethod digunakan untuk mengambil list-list metode pembayaran yg bisa dilakukan
	GetAvailablePaymentMethod(cart *model.Cart) error
}

// NewProcessor adalah sebuah konstruk untuk mengakses package 'Processor' ini
func NewProcessor(redisRepo repository.RedisRepositoryInterface, cartRepo repository.CartRepositoryInterface, courierRepo repository.CourierRepositoryInterface, userAddressRepo repository.UserAddressRepositoryInterface, paymentMethodRepo repository.PaymentMethodRepositoryInterface) ProcessorInterface {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return &processor{
		redisRepo:         redisRepo,
		cartRepo:          cartRepo,
		courierRepo:       courierRepo,
		userAddressRepo:   userAddressRepo,
		paymentMethodRepo: paymentMethodRepo,
		client:            client,
	}
}
