package stockrepo

import "github.com/pandudpn/shopping-cart/src/domain/model"

const (
	QUERY_SELECT              = "select id, product_id, quantity_hold, created_at, updated_at from stock "
	QUERY_UPDATE_STOCK        = "update stock set quantity_hold = quantity_hold + $1 where id = $2"
	QUERY_STOCK_BY_PRODUCT_ID = QUERY_SELECT + "where product_id=$1"
)

func (sr *StockRepository) FindStockByProductId(product *model.Product) (*model.Stock, error) {
	row := sr.DB.QueryRow(QUERY_STOCK_BY_PRODUCT_ID, product.Id)

	stock, err := rowToStock(row)
	if err != nil {
		return nil, err
	}

	stock.SetProduct(product)
	product.SetStock(stock)

	return stock, nil
}
