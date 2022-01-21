package orderdeliveryrepo

import (
	"errors"

	"github.com/pandudpn/shopping-cart/src/domain/model"
)

func (odr *OrderDeliveryRepository) UpdateOrderDelivery(order *model.Order) error {
	orderDelivery := order.GetDelivery()

	stmt, err := odr.DB.Prepare(QUERY_UPDATE)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(orderDelivery.Id, orderDelivery.Status, orderDelivery.DeliveredAt, orderDelivery.PackageReceivedAt, orderDelivery.UpdatedAt, orderDelivery.RefId)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		err = errors.New("update order_delivery failed")
		return err
	}

	return nil
}
