package cartproductrepo

import "github.com/pandudpn/shopping-cart/src/domain/model"

const (
	QUERY_SELECT = "select cp.id, cp.cart_id, cp.product_id, cp.quantity, cp.base_price, " +
		"cp.total_price, p.id, p.name, p.slug, pc.id, pc.name, p.price, p.discounted_price, " +
		"p.qty, pc.slug, s.id, coalesce(s.quantity_hold, 0), p.weight " +
		"from cart_product cp " +
		"inner join cart c on cp.cart_id = c.id " +
		"inner join product p on cp.product_id = p.id " +
		"inner join product_category pc on pc.id = p.category_id " +
		"left join stock s on p.id = s.product_id "

	QUERY_SELECT_IMAGE = "select mf.id, mf.filename, mf.url " +
		"from product_image pi inner join media_file mf on mf.id = pi.image_id " +
		"where pi.product_id = $1 and mf.deleted = false limit 1"

	QUERY_BY_CART_ID            = QUERY_SELECT + "where cart_id = $1 order by cp.created_at asc"
	QUERY_BY_CART_ID_PRODUCT_ID = QUERY_SELECT + "where cart_id = $1 and cp.product_id = $2"
)

func (cpr *CartProductRepository) FindCartProductsByCartId(cart *model.Cart) error {
	rows, err := cpr.DB.Query(QUERY_BY_CART_ID, cart.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		cartProduct, err := rowsToCartProduct(rows)
		if err != nil {
			return err
		}

		cart.AddProduct(cartProduct)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

func (cpr *CartProductRepository) FindCartProductByCartIdAndProductId(cartId, productId int) (*model.CartProduct, error) {
	row := cpr.DB.QueryRow(QUERY_BY_CART_ID_PRODUCT_ID, cartId, productId)

	cartProduct, err := rowToCartProduct(row)
	if err != nil {
		return nil, err
	}

	return cartProduct, nil
}
