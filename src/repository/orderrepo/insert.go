package orderrepo

import "github.com/pandudpn/shopping-cart/src/domain/model"

func (o *OrderRepository) CreateOrder(order *model.Order) error {
	stmt, err := o.DB.Prepare(QUERY_INSERT)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(order.CartId, order.UserId, order.DeliveryAddressId, order.OrderNumber, order.TotalProductsPrice, order.TotalDeliveryCost, order.TotalPayment, order.Status, order.CreatedAt).Scan(&order.Id)
	return err
}
