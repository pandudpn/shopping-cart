package stockrepo

import "errors"

func (sr *StockRepository) UpdateStock(stockId, qty int) error {
	stmt, err := sr.DB.Prepare(QUERY_UPDATE_STOCK)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(qty, stockId)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		err = errors.New("stock.not_updated")
		return err
	}

	return nil
}
