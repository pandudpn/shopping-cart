package orderproductrepo

import (
	"database/sql"

	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

type OrderProductRepository struct {
	DB dbc.SqlDbc
}

const (
	QUERY_INSERT = "insert into order_product (order_id, product_id, quantity, base_price, total_price, created_at) " +
		"values($1, $2, $3, $4, $5, $6) returning id"
	QUERY_SELECT = "select op.id, op.quantity, op.base_price, op.total_price, p.id, p.name, p.slug, p.price, c.id, c.name, c.slug, p.weight " +
		"from order_product as op inner join product as p on p.id = op.product_id " +
		"inner join product_category as c on c.id = p.category_id "
	QUERY_BY_ORDER_ID = QUERY_SELECT + "where op.order_id = $1"
)

func rowsToOrderProduct(rows *sql.Rows) (*model.OrderProduct, error) {
	orderProduct := &model.OrderProduct{}
	product := &model.Product{}
	category := &model.ProductCategory{}

	err := rows.Scan(
		&orderProduct.Id, &orderProduct.Quantity, &orderProduct.BasePrice, &orderProduct.TotalPrice,
		&product.Id, &product.Name, &product.Slug, &product.Price, &category.Id, &category.Name,
		&category.Slug, &product.Weight,
	)
	if err != nil {
		return nil, err
	}

	product.SetCategory(category)
	orderProduct.SetProduct(product)

	return orderProduct, nil
}
