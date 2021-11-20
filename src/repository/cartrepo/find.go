package cartrepo

import "github.com/pandudpn/shopping-cart/src/domain/model"

const (
	QUERY_SELECT = "select id, user_id, user_address_id, courier_id, payment_method_id, " +
		"is_active, key, created_at from cart "
	QUERY_ACTIVE_CART_BY_USER_ID = QUERY_SELECT + "where user_id = $1 and is_active = true"
	QUERY_CART_BY_KEY            = QUERY_SELECT + "where key = $1 and is_active = true"
)

func (cr *CartRepository) FindActiveCartByUserId(userId int) (*model.Cart, error) {
	row := cr.DB.QueryRow(QUERY_ACTIVE_CART_BY_USER_ID, userId)

	cart, err := rowToCart(row)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (cr *CartRepository) FindCartByKey(key string) (*model.Cart, error) {
	row := cr.DB.QueryRow(QUERY_CART_BY_KEY, key)

	cart, err := rowToCart(row)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
