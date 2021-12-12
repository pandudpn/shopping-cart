package orderdeliverystatusrepo

import (
	"database/sql"

	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

type OrderDeliveryStatusRepository struct {
	DB dbc.SqlDbc
}

const (
	QUERY_INSERT            = "insert into order_delivery_status (order_delivery_id, status, created_at) values ($1, $2, $3)"
	QUERY_SELECT            = "select id, status, created_at from order_delivery_status "
	QUERY_BY_ORDER_DELIVERY = QUERY_SELECT + "where order_delivery_id=$1"
)

func rowsToOrderDeliveryStatus(rows *sql.Rows) (*model.OrderDeliveryStatus, error) {
	orderDeliveryStatus := &model.OrderDeliveryStatus{}

	err := rows.Scan(&orderDeliveryStatus.Id, &orderDeliveryStatus.Status, &orderDeliveryStatus.CreatedAt)
	if err != nil {
		return nil, err
	}

	return orderDeliveryStatus, nil
}
