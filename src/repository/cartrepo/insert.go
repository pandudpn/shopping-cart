package cartrepo

import "github.com/pandudpn/shopping-cart/src/domain/model"

const (
	QUERY_INSERT = "insert into cart (user_id, is_active, key, created_at) " +
		"values ($1, $2, $3, $4) returning id"
)

func (cr *CartRepository) InsertNewCart(cart *model.Cart) error {
	stmt, err := cr.DB.Prepare(QUERY_INSERT)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(cart.UserId, cart.IsActive, cart.Key, cart.CreatedAt).Scan(&cart.Id)
	if err != nil {
		return err
	}

	return nil
}
