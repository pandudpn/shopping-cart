package orderrepo

import (
	"errors"
	"time"

	"github.com/pandudpn/shopping-cart/src/domain/model"
)

func (or *OrderRepository) UpdateOrder(order *model.Order) error {
	now := time.Now().UTC()

	stmt, err := or.DB.Prepare(QUERY_UPDATE)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(order.Status, order.CompletedAt, order.CanceledAt, now, order.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		err = errors.New("update order failed")
		return err
	}

	return nil
}
