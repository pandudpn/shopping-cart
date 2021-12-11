package orderpaymentrepo

import "github.com/pandudpn/shopping-cart/src/domain/model"

func (opr *OrderPaymentRepository) FindPaymentByOrder(order *model.Order) error {
	row := opr.DB.QueryRow(QUERY_BY_ORDER, order.Id)

	orderPayment, err := rowToOrderPayment(row)
	if err != nil {
		return err
	}

	order.SetPayment(orderPayment)
	return nil
}
