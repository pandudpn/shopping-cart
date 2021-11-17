package stockrepo

import (
	"database/sql"

	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

type StockRepository struct {
	DB dbc.SqlDbc
}

func rowToStock(row *sql.Row) (*model.Stock, error) {
	stock := &model.Stock{}

	err := row.Scan(&stock.Id, &stock.ProductId, &stock.QuantityHold, &stock.CreatedAt, &stock.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return stock, nil
}
