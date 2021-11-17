package manager

import (
	"database/sql"
	"errors"
	"reflect"

	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/repository"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

type cartManager struct {
	cartRepo        repository.CartRepositoryInterface
	cartProductRepo repository.CartProductRepositoryInterface
}

func NewCartManager(cart repository.CartRepositoryInterface) CartManagerInterface {
	return &cartManager{
		cartRepo: cart,
	}
}

func (cm *cartManager) GetActiveCart(key string, userId int) (*model.Cart, error) {
	activeCart := model.NewCart()

	if reflect.ValueOf(key).IsZero() {
		cart, err := cm.cartRepo.FindActiveCartByUserId(userId)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				logger.Log.Errorf("error query %v", err)
				err = errors.New("query.find.error")
				return nil, err
			}

			return activeCart, nil
		}

		activeCart = cart
	} else {
		cart, err := cm.cartRepo.FindCartByKey(key)
		if err != nil {
			logger.Log.Errorf("error query %v", err)
			err = errors.New("query.find.error")
			return nil, err
		}

		if cart.UserId != userId {
			err = errors.New("cart.not_yours")
			return nil, err
		}

		activeCart = cart
	}

	// query for get all relation cart will be here
	err := cm.cartProductRepo.FindCartProductsByCartId(activeCart)
	if err != nil {
		logger.Log.Errorf("error get products query %v", err)
		err = errors.New("query.find.error")
		return nil, err
	}

	return activeCart, nil
}
