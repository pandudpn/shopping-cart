package cartrepo

import (
	"errors"

	"github.com/pandudpn/shopping-cart/src/domain/model"
)

const (
	QUERY_UPDATE = "update cart set courier_id=$1, user_address_id=$2, payment_method_id=$3, updated_at=now() where id=$4"
)

func (cr *CartRepository) UpdateCart(cart *model.Cart) error {
	stmt, err := cr.DB.Prepare(QUERY_UPDATE)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(cart.CourierId, cart.UserAddressId, cart.PaymentMethodId, cart.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		err = errors.New("update.failed")
		return err
	}

	return nil
}
