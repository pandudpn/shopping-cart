package cartrepo

import (
	"database/sql"

	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

type CartRepository struct {
	DB dbc.SqlDbc
}

func rowToCart(row *sql.Row) (*model.Cart, error) {
	cart := &model.Cart{}

	err := row.Scan(&cart.Id, &cart.UserId, &cart.UserAddressId, &cart.CourierId, &cart.PaymentMethodId, &cart.IsActive, &cart.Key, &cart.CreatedAt)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
