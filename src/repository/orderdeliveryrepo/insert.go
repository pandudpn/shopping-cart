package orderdeliveryrepo

import "github.com/pandudpn/shopping-cart/src/domain/model"

func (odr *OrderDeliveryRepository) CreateOrderDelivery(orderDelivery *model.OrderDelivery) error {
	stmt, err := odr.DB.Prepare(QUERY_INSERT)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		orderDelivery.OrderId, orderDelivery.CourierId, orderDelivery.DeliveryAddressId, orderDelivery.DeliveryCost,
		orderDelivery.DeliveryCostDiscount, orderDelivery.TotalDeliveryCost, orderDelivery.TrackingNumber, orderDelivery.Address,
		orderDelivery.ReceivedName, orderDelivery.PhoneNumber, orderDelivery.Lat, orderDelivery.Long, orderDelivery.Status, orderDelivery.CreatedAt,
	).Scan(&orderDelivery.Id)
	return err
}
