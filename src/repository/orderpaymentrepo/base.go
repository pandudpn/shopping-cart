package orderpaymentrepo

import (
	"database/sql"

	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

type OrderPaymentRepository struct {
	DB dbc.SqlDbc
}

const (
	QUERY_INSERT = "insert into order_payment (order_id, payment_method_id, total_payment, total_paid, status, is_active, qr_link, " +
		"redirect_link, deep_link, va_number, expired_at, created_at) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) returning id"
	QUERY_SELECT = "select op.id, op.total_payment, op.total_paid, op.status, op.qr_link, op.redirect_link, op.deep_link, " +
		"op.va_number, op.expired_at, pm.id, pm.name, pm.category, pm.image " +
		"from order_payment as op inner join payment_method as pm on pm.id = op.payment_method_id "
	QUERY_BY_ORDER = QUERY_SELECT + "where op.order_id = $1"
)

func rowToOrderPayment(row *sql.Row) (*model.OrderPayment, error) {
	orderPayment := &model.OrderPayment{}
	paymentMethod := &model.PaymentMethod{}

	err := row.Scan(
		&orderPayment.Id, &orderPayment.TotalPayment, &orderPayment.TotalPaid, &orderPayment.Status, &orderPayment.QrLink,
		&orderPayment.RedirectLink, &orderPayment.DeepLink, &orderPayment.VaNumber, &orderPayment.ExpiredAt,
		&paymentMethod.Id, &paymentMethod.Name, &paymentMethod.Category, &paymentMethod.Image,
	)
	if err != nil {
		return nil, err
	}

	orderPayment.SetPaymentMethod(paymentMethod)
	return orderPayment, nil
}
