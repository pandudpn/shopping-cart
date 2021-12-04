package courierrepo

import "github.com/pandudpn/shopping-cart/src/domain/model"

func (cr *CourierRepository) FindEnabledCourier() ([]*model.Courier, error) {
	couriers := make([]*model.Courier, 0)

	rows, err := cr.DB.Query(QUERY_SELECT)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		courier, err := rowsToCourier(rows)
		if err != nil {
			return nil, err
		}

		couriers = append(couriers, courier)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return couriers, nil
}

func (cr *CourierRepository) FindCourierById(courierId int) (*model.Courier, error) {
	row := cr.DB.QueryRow(QUERY_BY_ID, courierId)

	courier, err := rowToCourier(row)
	return courier, err
}
