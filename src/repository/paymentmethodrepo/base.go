package paymentmethodrepo

import (
	"database/sql"

	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

const (
	QUERY_SELECT = "select pm.id, pm.code, pm.category, pm.name, pm.image, pm.key " +
		"from payment_method pm where pm.enabled = true "
	QUERY_BY_ID = QUERY_SELECT + "and pm.id = $1"
)

type PaymentMethodRepository struct {
	DB dbc.SqlDbc
}

func rowToPaymentMethod(row *sql.Row) (*model.PaymentMethod, error) {
	paymentMethod := &model.PaymentMethod{}

	err := row.Scan(&paymentMethod.Id, &paymentMethod.Code, &paymentMethod.Category, &paymentMethod.Name, &paymentMethod.Image, &paymentMethod.Key)
	if err != nil {
		return nil, err
	}

	return paymentMethod, nil
}

func rowsToPaymentMethod(rows *sql.Rows) (*model.PaymentMethod, error) {
	paymentMethod := &model.PaymentMethod{}

	err := rows.Scan(&paymentMethod.Id, &paymentMethod.Code, &paymentMethod.Category, &paymentMethod.Name, &paymentMethod.Image, &paymentMethod.Key)
	if err != nil {
		return nil, err
	}

	return paymentMethod, nil
}
