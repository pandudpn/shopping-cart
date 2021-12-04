package cartusecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/pandudpn/shopping-cart/src/api/presenter/cartpresenter"
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (cu *CartUseCase) AddToCart(ctx context.Context, req *model.RequestAddToCart) utils.ResponseInterface {
	var (
		cart        = &model.Cart{}
		err         error
		isAddToCart = true
	)

	err = cu.TxRepo.TxEnd(func() error {
		activeCart, err := cu.addToCart(req)
		if err != nil {
			return err
		}

		cart = activeCart
		return nil
	})

	if err != nil {
		return cartpresenter.ResponseCart(isAddToCart, nil, err)
	}

	return cartpresenter.ResponseCart(isAddToCart, cart, nil)
}

func (cu *CartUseCase) GetCart(ctx context.Context, userId int, key string) utils.ResponseInterface {
	var err error

	cart, err := cu.getActiveCart(key, userId)
	if err != nil {
		logger.Log.Errorf("error get cart %v", err)
		return cartpresenter.ResponseCart(false, nil, err)
	}

	return cartpresenter.ResponseCart(false, cart, nil)
}

func (cu *CartUseCase) addToCart(req *model.RequestAddToCart) (*model.Cart, error) {
	product, err := cu.ProductRepo.FindProductById(req.ProductId)
	if err != nil {
		logger.Log.Errorf("error get product %v", err)
		return nil, errProductNotFound
	}

	if product.GetQuantity() < req.Quantity {
		logger.Log.Errorf("stock on database insufficient %d and from request %d", product.GetQuantity(), req.Quantity)
		return nil, errQuantity
	}

	cart, err := cu.getActiveCart(req.SecretKey, req.UserId)
	if err != nil {
		logger.Log.Errorf("error get active cart %v", err)
		return nil, errActiveCart
	}

	cartProduct, err := cu.CartProductRepo.FindCartProductByCartIdAndProductId(cart.Id, product.Id)
	if err != nil {
		logger.Log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			newCartProduct := model.NewCartProduct()
			newCartProduct.CartId = cart.Id
			newCartProduct.ProductId = product.Id
			newCartProduct.BasePrice = product.Price
			newCartProduct.Quantity = req.Quantity
			newCartProduct.TotalPrice = product.Price * float64(req.Quantity)
			newCartProduct.SetCart(cart)
			newCartProduct.SetProduct(product)

			err = cu.CartProductRepo.InsertNewCartProduct(newCartProduct)
			if err != nil {
				logger.Log.Errorf("error insert product %v", err)
				return nil, errQueryInsert
			}

			cart.AddProduct(newCartProduct)
			return cart, nil
		}

		err = errors.New("query.find.error")
		return nil, err
	}
	// remove terlebih dahulu, baru nanti di add kembali
	cart.RemoveProduct(cartProduct)

	cartProduct.Quantity = req.Quantity
	cartProduct.TotalPrice = float64(req.Quantity) * product.Price

	err = cu.CartProductRepo.UpdateCartProduct(cartProduct)
	if err != nil {
		logger.Log.Error(err)
		err = errors.New("query.update.failed")

		return nil, err
	}
	cart.AddProduct(cartProduct)

	return cart, nil
}
