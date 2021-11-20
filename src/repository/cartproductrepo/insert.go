package cartproductrepo

import "github.com/pandudpn/shopping-cart/src/domain/model"

const (
	QUERY_INSERT = "insert into cart_product (cart_id, product_id, base_price, quantity, total_price, created_at) " +
		"values($1, $2, $3, $4, $5, $6) returning id"
)

func (cpr *CartProductRepository) InsertNewCartProduct(cartProduct *model.CartProduct) error {
	stmt, err := cpr.DB.Prepare(QUERY_INSERT)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		cartProduct.CartId, cartProduct.ProductId, cartProduct.BasePrice,
		cartProduct.Quantity, cartProduct.TotalPrice, cartProduct.CreatedAt,
	).Scan(&cartProduct.Id)
	if err != nil {
		return err
	}

	return nil
}
