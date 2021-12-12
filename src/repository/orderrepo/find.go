package orderrepo

import "github.com/pandudpn/shopping-cart/src/domain/model"

func (or *OrderRepository) FindOrderByOrderNumber(orderNumber string) (*model.Order, error) {
	row := or.DB.QueryRow(QUERY_BY_ORDER_NUMBER, orderNumber)

	order, err := rowToOrder(row)
	return order, err
}
