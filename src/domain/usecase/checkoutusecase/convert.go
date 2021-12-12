package checkoutusecase

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"

	"github.com/pandudpn/shopping-cart/src/domain/model"
	paymenthandler "github.com/pandudpn/shopping-cart/src/payment_handler"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (cu *CheckoutUseCase) convertCartToOrder(cart *model.Cart) (*model.Order, error) {
	var (
		err     error
		wg      sync.WaitGroup
		errChan = make(chan error, 1)
	)

	max := big.NewInt(99999999)
	n, _ := rand.Int(rand.Reader, max)
	orderNumber := fmt.Sprintf("%08d", n.Int64())

	cart.SetOrderNumber(orderNumber)

	handler, err := paymenthandler.GetHandlerPayment(cart.GetPaymentMethod().Code).Process(cart)
	if err != nil {
		logger.Log.Errorf("error payment handler %v", err)
		return nil, errCreatePayment
	}
	orderPayment := handler.(*model.OrderPayment)

	order := cu.createOrder(cart)
	err = cu.OrderRepo.CreateOrder(order)
	if err != nil {
		logger.Log.Errorf("error create order %v", err)
		return nil, errInsert
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		err = cu.createOrderPayment(orderPayment, order)
		if err != nil {
			errChan <- err
			return
		}

		err = cu.createOrderDelivery(order, cart)
		if err != nil {
			errChan <- err
			return
		}

		err = cu.createOrderProducts(order, cart)
		if err != nil {
			errChan <- err
			return
		}
	}()
	select {
	case err = <-errChan:
		return nil, errInsert
	default:
	}
	wg.Wait()

	cart.IsActive = false
	err = cu.CartRepo.UpdateCart(cart)
	if err != nil {
		logger.Log.Errorf("error update cart to inactive %v", err)
		return nil, errUpdate
	}

	return order, nil
}

func (cu *CheckoutUseCase) createOrder(cart *model.Cart) *model.Order {
	order := model.NewOrder()
	order.CartId = cart.Id
	order.DeliveryAddressId = *cart.UserAddressId
	order.TotalDeliveryCost = cart.TotalDeliveryCost
	order.TotalPayment = cart.Total
	order.TotalProductsPrice = cart.TotalProductsPrice
	order.UserId = cart.UserId
	order.OrderNumber = cart.GetOrderNumber()

	return order
}

func (cu *CheckoutUseCase) createOrderPayment(orderPayment *model.OrderPayment, order *model.Order) error {
	orderPayment.OrderId = order.Id

	err := cu.OrderPaymentRepo.CreateOrderPayment(orderPayment)
	if err != nil {
		logger.Log.Errorf("error create order payment %v", err)
		return err
	}

	order.SetPayment(orderPayment)
	return nil
}

func (cu *CheckoutUseCase) createOrderDelivery(order *model.Order, cart *model.Cart) error {
	orderDelivery := model.NewOrderDelivery()
	orderDelivery.CourierId = cart.GetCourier().Id
	orderDelivery.OrderId = order.Id
	orderDelivery.DeliveryAddressId = cart.GetUserAddress().Id
	orderDelivery.DeliveryCost = cart.GetCourier().DeliveryCost
	orderDelivery.DeliveryCostDiscount = cart.GetCourier().DeliveryCostDiscount
	orderDelivery.TotalDeliveryCost = cart.GetCourier().TotalDeliveryCost
	orderDelivery.Address = cart.GetUserAddress().Address
	orderDelivery.ReceivedName = cart.GetUserAddress().ReceiverName
	orderDelivery.PhoneNumber = cart.GetUserAddress().PhoneNumber
	orderDelivery.Lat = cart.GetUserAddress().Lat
	orderDelivery.Long = cart.GetUserAddress().Long
	orderDelivery.Rate = &cart.GetCourier().Rate

	err := cu.OrderDeliveryRepo.CreateOrderDelivery(orderDelivery)
	if err != nil {
		logger.Log.Errorf("error create order delivery %v", err)
		return err
	}

	orderDelivery.SetCourier(cart.GetCourier())
	orderDelivery.SetUserDelivery(cart.GetUserAddress())
	order.SetDelivery(orderDelivery)
	return nil
}

func (cu *CheckoutUseCase) createOrderProducts(order *model.Order, cart *model.Cart) error {
	for _, cartProduct := range cart.GetProducts() {
		orderProduct := model.NewOrderProduct()
		orderProduct.OrderId = order.Id
		orderProduct.BasePrice = cartProduct.BasePrice
		orderProduct.TotalPrice = cartProduct.TotalPrice
		orderProduct.Quantity = cartProduct.Quantity
		orderProduct.ProductId = cartProduct.ProductId

		err := cu.OrderProductRepo.CreateOrderProduct(orderProduct)
		if err != nil {
			logger.Log.Errorf("error create order product %v", err)
			return err
		}

		orderProduct.SetProduct(cartProduct.GetProduct())
		order.AddProduct(orderProduct)
	}

	return nil
}
