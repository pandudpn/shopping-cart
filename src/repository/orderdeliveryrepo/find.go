package orderdeliveryrepo

import "github.com/pandudpn/shopping-cart/src/domain/model"

func (odr *OrderDeliveryRepository) FindDeliveryByOrder(order *model.Order) error {
	row := odr.DB.QueryRow(QUERY_BY_ORDER, order.Id)

	orderDelivery, err := rowToOrderDelivery(row)
	if err != nil {
		return err
	}

	order.SetDelivery(orderDelivery)
	return nil
}
