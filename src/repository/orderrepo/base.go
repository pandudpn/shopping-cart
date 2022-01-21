package orderrepo

import (
	"database/sql"

	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

const (
	QUERY_INSERT = "insert into public.order (cart_id, user_id, delivery_address_id, " +
		"order_number, total_products_price, total_delivery_cost, total_payment, status, " +
		"created_at) " +
		"values($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id"
	QUERY_UPDATE = "update public.order set status=$1, completed_at=$2, canceled_at=$3, updated_at=$4 where id=$5"
	QUERY_SELECT = "select o.id, o.user_id, o.delivery_address_id, o.order_number, o.total_products_price, " +
		"o.total_delivery_cost, o.total_payment, o.status, o.created_at " +
		"from public.order AS o "
	QUERY_BY_ORDER_NUMBER = QUERY_SELECT + "where order_number=$1"
)

type OrderRepository struct {
	DB dbc.SqlDbc
}

func rowToOrder(row *sql.Row) (*model.Order, error) {
	order := &model.Order{}

	err := row.Scan(&order.Id, &order.UserId, &order.DeliveryAddressId, &order.OrderNumber, &order.TotalProductsPrice, &order.TotalDeliveryCost, &order.TotalPayment, &order.Status, &order.CreatedAt)
	return order, err
}
