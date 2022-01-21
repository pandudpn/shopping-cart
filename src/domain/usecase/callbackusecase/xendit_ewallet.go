package callbackusecase

import (
	"context"
	"time"

	"github.com/pandudpn/shopping-cart/src/api/presenter/callbackpresenter"
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (cu *CallbackUseCase) CallbackEWallet(ctx context.Context, req *model.EWalletPaymentNotification) utils.ResponseInterface {
	dataReq := req.Data

	order, err := cu.OrderRepo.FindOrderByOrderNumber(dataReq.ReferenceId)
	if err != nil {
		logger.Log.Errorf("error find order number %v. get order number %s", err, dataReq.ReferenceId)
		return callbackpresenter.ResponseCallback(nil, errOrderNotFound)
	}

	err = cu.hydrator(order)
	if err != nil {
		return callbackpresenter.ResponseCallback(nil, err)
	}

	err = cu.TxRepo.TxEnd(func() error {
		err = cu.checkStatusPayment(ctx, order, dataReq)
		if err != nil {
			return err
		}

		if order.GetPayment().Status == model.StatusPaymentSuccess {
			err = cu.createOrderCourier(order)
			return err
		}

		return nil
	})
	return callbackpresenter.ResponseCallback(order, err)
}

func (cu *CallbackUseCase) checkStatusPayment(ctx context.Context, order *model.Order, data *model.DataEWalletNotification) error {
	now := time.Now().UTC()

	order.UpdatedAt = &now
	order.GetPayment().UpdatedAt = &now

	if !data.IsStatusSuccess() {
		if !data.IsStatusRefund() {
			order.CanceledAt = &now
			order.SetStatusToFailed()
			order.GetPayment().SetStatusToFailed()
		} else {
			order.SetStatusToReturn()
			order.GetPayment().SetStatusToRefund()
		}
	} else {
		order.SetStatusToPaymentConfirmed()
		order.GetPayment().SetStatusToSuccess()
		order.GetPayment().ConfirmedAt = &now
	}

	err := cu.OrderRepo.UpdateOrder(order)
	if err != nil {
		logger.Log.Errorf("error update data order %v", err)
		return errUpdateOrder
	}

	err = cu.OrderPaymentRepo.UpdateOrderPayment(order)
	if err != nil {
		logger.Log.Errorf("error update data order payment %v", err)
		return errUpdateOrder
	}

	return nil
}
