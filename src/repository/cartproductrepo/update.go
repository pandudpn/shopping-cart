package cartproductrepo

import (
	"errors"

	"github.com/pandudpn/shopping-cart/src/domain/model"
)

const (
	QUERY_UPDATE = "update cart_product set quantity=$1, total_price=$2 where id=$3"
)

func (cpr *CartProductRepository) UpdateCartProduct(cartProduct *model.CartProduct) error {
	stmt, err := cpr.DB.Prepare(QUERY_UPDATE)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(cartProduct.Quantity, cartProduct.TotalPrice, cartProduct.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		err = errors.New("query.update.failed")
		return err
	}

	return nil
}
