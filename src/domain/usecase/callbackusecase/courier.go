package callbackusecase

import (
	courierhandler "github.com/pandudpn/shopping-cart/src/courier_handler"
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (cu *CallbackUseCase) createOrderCourier(order *model.Order) error {
	err := courierhandler.GetCourierHandler(courierhandler.Shipper).Process(order)
	if err != nil {
		return err
	}

	err = cu.OrderDeliveryRepo.UpdateOrderDelivery(order)
	if err != nil {
		logger.Log.Errorf("error update order_delivery %v", err)
		return errUpdateOrder
	}

	return nil
}
