package cartproductrepo

import (
	"database/sql"

	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

type CartProductRepository struct {
	DB dbc.SqlDbc
}

func rowsToCartProduct(rows *sql.Rows) (*model.CartProduct, error) {
	cartProduct := &model.CartProduct{}
	product := &model.Product{}
	category := &model.ProductCategory{}
	stock := &model.Stock{}

	err := rows.Scan(
		&cartProduct.Id, &cartProduct.CartId, &cartProduct.ProductId, &cartProduct.Quantity,
		&cartProduct.BasePrice, &cartProduct.TotalPrice, &product.Id, &product.Name,
		&product.Slug, &category.Id, &category.Name, &product.Price, &product.DiscountedPrice,
		&product.Qty, &category.Slug, &stock.Id, &stock.QuantityHold,
	)

	if err != nil {
		return nil, err
	}

	product.SetCategory(category)
	product.SetStock(stock)
	cartProduct.SetProduct(product)

	return cartProduct, nil
}

func rowToCartProduct(row *sql.Row) (*model.CartProduct, error) {
	cartProduct := &model.CartProduct{}
	product := &model.Product{}
	category := &model.ProductCategory{}
	stock := &model.Stock{}

	err := row.Scan(
		&cartProduct.Id, &cartProduct.CartId, &cartProduct.ProductId, &cartProduct.Quantity,
		&cartProduct.BasePrice, &cartProduct.TotalPrice, &product.Id, &product.Name,
		&product.Slug, &category.Id, &category.Name, &product.Price, &product.DiscountedPrice,
		&product.Qty, &category.Slug, &stock.Id, &stock.QuantityHold,
	)

	if err != nil {
		return nil, err
	}

	product.SetCategory(category)
	product.SetStock(stock)
	cartProduct.SetProduct(product)

	return cartProduct, nil
}

func rowToImage(row *sql.Row) (*model.ProductImage, error) {
	productImage := &model.ProductImage{}
	mediaFile := &model.MediaFile{}

	err := row.Scan(&mediaFile.Id, &mediaFile.Filename, &mediaFile.Url)
	if err != nil {
		return nil, err
	}

	productImage.SetImage(mediaFile)
	return productImage, nil
}
