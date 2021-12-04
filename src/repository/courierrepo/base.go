package courierrepo

import (
	"database/sql"

	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

const (
	QUERY_SELECT = "select c.id, c.code, c.name, c.category, c.image, c.enabled " +
		"from courier c where c.enabled = true "
	QUERY_BY_ID = QUERY_SELECT + "and c.id = $1"
)

type CourierRepository struct {
	DB dbc.SqlDbc
}

func rowToCourier(row *sql.Row) (*model.Courier, error) {
	courier := &model.Courier{}

	err := row.Scan(&courier.Id, &courier.Code, &courier.Name, &courier.Category, &courier.Image, &courier.Enabled)
	if err != nil {
		return nil, err
	}

	return courier, nil
}

func rowsToCourier(rows *sql.Rows) (*model.Courier, error) {
	courier := &model.Courier{}

	err := rows.Scan(&courier.Id, &courier.Code, &courier.Name, &courier.Category, &courier.Image, &courier.Enabled)
	if err != nil {
		return nil, err
	}

	return courier, nil
}
