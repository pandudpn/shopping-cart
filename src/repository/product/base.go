package product

import (
	"database/sql"

	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

type ProductRepository struct {
	DB dbc.SqlDbc
}

func rowsToProduct(rows *sql.Rows) (*model.Product, error) {
	product := &model.Product{}
	category := &model.ProductCategory{}

	err := rows.Scan(&product.Id, &product.Name, &product.Slug, &product.Description, &product.Price, &product.DiscountedPrice, &product.Qty, &product.Enabled, &product.CreatedAt, &category.Id, &category.Name, &category.Slug)
	if err != nil {
		return nil, err
	}

	product.SetCategory(category)
	return product, nil
}

func rowToProduct(row *sql.Row) (*model.Product, error) {
	product := &model.Product{}
	category := &model.ProductCategory{}

	err := row.Scan(&product.Id, &product.Name, &product.Slug, &product.Description, &product.Price, &product.DiscountedPrice, &product.Qty, &product.Enabled, &product.CreatedAt, &category.Id, &category.Name, &category.Slug)
	if err != nil {
		return nil, err
	}

	product.SetCategory(category)
	return product, nil
}
