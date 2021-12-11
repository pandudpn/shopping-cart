package orderdeliverystatusrepo

import (
	"errors"

	"github.com/pandudpn/shopping-cart/src/domain/model"
)

func (odsr *OrderDeliveryStatusRepository) CreateOrderDeliveryStatus(orderDeliveryStatus *model.OrderDeliveryStatus) error {
	stmt, err := odsr.DB.Prepare(QUERY_INSERT)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(orderDeliveryStatus.OrderDeliveryId, orderDeliveryStatus.Status, orderDeliveryStatus.CreatedAt)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		err = errors.New("failed create order delivery status")
		return err
	}

	return nil
}
