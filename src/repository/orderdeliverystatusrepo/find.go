package orderdeliverystatusrepo

import "github.com/pandudpn/shopping-cart/src/domain/model"

func (odsr *OrderDeliveryStatusRepository) FindStatusByOrderDelivery(orderDelivery *model.OrderDelivery) error {
	statuses := make([]*model.OrderDeliveryStatus, 0)

	rows, err := odsr.DB.Query(QUERY_BY_ORDER_DELIVERY, orderDelivery.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		status, err := rowsToOrderDeliveryStatus(rows)
		if err != nil {
			return err
		}

		statuses = append(statuses, status)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	orderDelivery.SetStatuses(statuses)
	return nil
}
