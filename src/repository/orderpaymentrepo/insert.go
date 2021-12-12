package orderpaymentrepo

import "github.com/pandudpn/shopping-cart/src/domain/model"

func (opr *OrderPaymentRepository) CreateOrderPayment(orderPayment *model.OrderPayment) error {
	stmt, err := opr.DB.Prepare(QUERY_INSERT)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		orderPayment.OrderId, orderPayment.PaymentMethodId, orderPayment.TotalPayment, orderPayment.TotalPaid,
		orderPayment.Status, orderPayment.IsActive, orderPayment.QrLink, orderPayment.RedirectLink, orderPayment.DeepLink,
		orderPayment.VaNumber, orderPayment.ExpiredAt, orderPayment.CreatedAt,
	).Scan(&orderPayment.Id)
	return err
}
