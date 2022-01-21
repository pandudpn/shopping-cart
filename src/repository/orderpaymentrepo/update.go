package orderpaymentrepo

import (
	"errors"

	"github.com/pandudpn/shopping-cart/src/domain/model"
)

func (opr *OrderPaymentRepository) UpdateOrderPayment(order *model.Order) error {
	orderPayment := order.GetPayment()

	stmt, err := opr.DB.Prepare(QUERY_UPDATE)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(orderPayment.Id, orderPayment.Status, orderPayment.ConfirmedAt, orderPayment.UpdatedAt)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		err = errors.New("update order_payment failed")
		return err
	}

	return nil
}
