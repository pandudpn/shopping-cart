package orderproductrepo

import "github.com/pandudpn/shopping-cart/src/domain/model"

func (opr *OrderProductRepository) CreateOrderProduct(orderProduct *model.OrderProduct) error {
	stmt, err := opr.DB.Prepare(QUERY_INSERT)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(orderProduct.OrderId, orderProduct.ProductId, orderProduct.Quantity, orderProduct.BasePrice, orderProduct.TotalPrice, orderProduct.CreatedAt).Scan(&orderProduct)
	return err
}
