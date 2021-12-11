package orderproductrepo

import "github.com/pandudpn/shopping-cart/src/domain/model"

func (opr *OrderProductRepository) FindProductsByOrder(order *model.Order) ([]*model.OrderProduct, error) {
	orderProducts := make([]*model.OrderProduct, 0)

	rows, err := opr.DB.Query(QUERY_BY_ORDER_ID, order.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		orderProduct, err := rowsToOrderProduct(rows)
		if err != nil {
			return nil, err
		}

		orderProducts = append(orderProducts, orderProduct)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orderProducts, nil
}
