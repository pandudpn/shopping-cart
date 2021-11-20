package cartproductrepo

import "github.com/pandudpn/shopping-cart/src/domain/model"

const (
	QUERY_SELECT = "select cp.id, cp.cart_id, cp.product_id, cp.quantity, cp.base_price, " +
		"cp.total_price, p.id, p.name, p.slug, pc.id, pc.name, pc.slug, s.id, coalesce(s.quantity_hold, 0) " +
		"from cart_product cp " +
		"inner join cart c on cp.cart_id = c.id " +
		"inner join product p on cp.product_id = p.id " +
		"inner join product_category pc on pc.id = p.category_id " +
		"left join stock s on p.id = s.product_id "

	QUERY_SELECT_IMAGE = "select mf.id, mf.filename, mf.url " +
		"from product_image pi inner join media_file mf on mf.id = pi.image_id " +
		"where pi.product_id = $1 and mf.deleted = false limit 1"

	QUERY_BY_CART_ID = QUERY_SELECT + "where cart_id = $1 order by cp.created_at asc"
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

		// query for get product image
		row := cpr.DB.QueryRow(QUERY_SELECT_IMAGE, cartProduct.ProductId)
		productImage, err := rowToImage(row)
		if err != nil {
			return err
		}

		cartProduct.GetProduct().AddImage(productImage)
		cart.AddProduct(cartProduct)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}
