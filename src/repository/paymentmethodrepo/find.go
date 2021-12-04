package paymentmethodrepo

import "github.com/pandudpn/shopping-cart/src/domain/model"

func (pmr *PaymentMethodRepository) FindEnabledPaymentMethod() ([]*model.PaymentMethod, error) {
	paymentMethods := make([]*model.PaymentMethod, 0)

	rows, err := pmr.DB.Query(QUERY_SELECT)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		paymentMethod, err := rowsToPaymentMethod(rows)
		if err != nil {
			return nil, err
		}

		paymentMethods = append(paymentMethods, paymentMethod)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return paymentMethods, nil
}

func (pmr *PaymentMethodRepository) FindPaymentMethodById(paymentMethodId int) (*model.PaymentMethod, error) {
	row := pmr.DB.QueryRow(QUERY_BY_ID, paymentMethodId)

	paymentMethod, err := rowToPaymentMethod(row)
	return paymentMethod, err
}
