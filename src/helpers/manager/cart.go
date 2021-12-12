package manager

import (
	"database/sql"
	"errors"
	"reflect"
	"sync"

	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/repository"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

type cartManager struct {
	cartRepo          repository.CartRepositoryInterface
	cartProductRepo   repository.CartProductRepositoryInterface
	imageRepo         repository.ProductImageRepositoryInterface
	userRepo          repository.UserRepositoryInterface
	userAddressRepo   repository.UserAddressRepositoryInterface
	courierRepo       repository.CourierRepositoryInterface
	paymentMethodRepo repository.PaymentMethodRepositoryInterface
}

func NewCartManager(
	cart repository.CartRepositoryInterface, cartProduct repository.CartProductRepositoryInterface, imageRepo repository.ProductImageRepositoryInterface, userRepo repository.UserRepositoryInterface,
	userAddressRepo repository.UserAddressRepositoryInterface, courierRepo repository.CourierRepositoryInterface, paymentMethodRepo repository.PaymentMethodRepositoryInterface,
) CartManagerInterface {
	return &cartManager{
		cartRepo:          cart,
		cartProductRepo:   cartProduct,
		imageRepo:         imageRepo,
		userRepo:          userRepo,
		userAddressRepo:   userAddressRepo,
		courierRepo:       courierRepo,
		paymentMethodRepo: paymentMethodRepo,
	}
}

func (cm *cartManager) GetActiveCart(key string, userId int, isCheckout bool) (*model.Cart, error) {
	activeCart := model.NewCart()

	if reflect.ValueOf(key).IsZero() {
		cart, err := cm.cartRepo.FindActiveCartByUserId(userId)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				logger.Log.Errorf("error query %v", err)
				err = errors.New("query.find.error")
				return nil, err
			}
		} else {
			activeCart = cart
		}
	} else {
		cart, err := cm.cartRepo.FindCartByKey(key)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				logger.Log.Errorf("error query %v", err)
				err = errors.New("query.find.error")
				return nil, err
			}
		} else {
			if cart.UserId != userId {
				err = errors.New("cart.not_yours")
				return nil, err
			}

			activeCart = cart
		}
	}

	wg := sync.WaitGroup{}
	// cek apakah cart baru atau cart existing
	logger.Log.Debug(activeCart.Id)
	if activeCart.Id == 0 { // cart baru
		activeCart.UserId = userId
	} else {
		logger.Log.Debugf("active cart %v", activeCart)
		// query for get all relation cart will be here
		err := cm.cartProductRepo.FindCartProductsByCartId(activeCart)
		if err != nil {
			logger.Log.Errorf("error get products query %v", err)
			err = errors.New("query.find.error")
			return nil, err
		}

		for _, cartProduct := range activeCart.GetProducts() {
			wg.Add(1)

			go func(cartProduct *model.CartProduct) {
				defer wg.Done()
				productImages, err := cm.imageRepo.FindImagesByProductId(cartProduct.ProductId)
				if err != nil {
					logger.Log.Errorf("error get images for product %d", cartProduct.ProductId)
					return
				}

				cartProduct.GetProduct().SetImages(productImages)
				activeCart.TotalProductsPrice += cartProduct.TotalPrice
			}(cartProduct)
		}
	}

	user, err := cm.userRepo.FindById(activeCart.UserId)
	if err != nil {
		logger.Log.Errorf("error get user from cart %v", err)
		err = errors.New("cart.user.not_found")
		return nil, err
	}

	activeCart.SetUser(user)
	if activeCart.UserAddressId != nil {
		userAddress, err := cm.userAddressRepo.FindUserAddressById(*activeCart.UserAddressId)
		if err != nil {
			logger.Log.Errorf("error get user_address from cart %v", err)
			err = errors.New("cart.user_address.not_found")
			return nil, err
		}

		activeCart.SetUserAddress(userAddress)
	}

	if activeCart.CourierId != nil {
		courier, err := cm.courierRepo.FindCourierById(*activeCart.CourierId)
		if err != nil {
			logger.Log.Errorf("error get courier from cart %v", err)
			err = errors.New("cart.courier.not_found")
			return nil, err
		}

		activeCart.SetCourier(courier)
	}

	if activeCart.PaymentMethodId != nil {
		paymentMethod, err := cm.paymentMethodRepo.FindPaymentMethodById(*activeCart.PaymentMethodId)
		if err != nil {
			logger.Log.Errorf("error get payment method from cart %v", err)
			err = errors.New("cart.payment_method.not_found")
			return nil, err
		}

		activeCart.SetPaymentMethod(paymentMethod)
	}

	wg.Wait()

	return activeCart, nil
}
