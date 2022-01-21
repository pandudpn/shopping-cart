package orderdeliveryrepo

import (
	"database/sql"
	
	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

type OrderDeliveryRepository struct {
	DB dbc.SqlDbc
}

const (
	QUERY_INSERT = "insert into order_delivery (order_id, courier_id, delivery_address_id, delivery_cost, delivery_cost_discount, " +
		"total_delivery_cost, tracking_number, address, receiver_name, phone_number, lat, long, status, created_at, rate) " +
		"values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) returning id"
	QUERY_UPDATE = "update order_delivery set status=$2, delivered_at=$3, package_received_at=$4, updated_at=$5, ref_id=$6 where id=$1"
	QUERY_SELECT = "select od.id, od.delivery_cost, od.delivery_cost_discount, od.total_delivery_cost, od.tracking_number, od.address, " +
		"od.receiver_name, od.phone_number, od.lat, od.long, od.status, od.delivered_at, od.package_received_at, c.id, c.name, c.category, c.image, " +
		"ua.post_code, pr.id, pr.name, ct.id, ct.name, dt.id, dt.name, a.id, a.name, a.lat, a.long, a.ref_id, od.rate, od.ref_id " +
		"from order_delivery as od inner join courier as c on od.courier_id = c.id " +
		"inner join user_address as ua on ua.id = od.delivery_address_id " +
		"inner join region as pr on pr.id = ua.province_id inner join region as ct on ct.id = ua.city_id " +
		"inner join region as dt on dt.id = ua.district_id inner join region as a on a.id = ua.area_id "
	QUERY_BY_ORDER = QUERY_SELECT + "where od.order_id = $1"
)

func rowToOrderDelivery(row *sql.Row) (*model.OrderDelivery, error) {
	orderDelivery := &model.OrderDelivery{}
	courier := &model.Courier{}
	userAddress := &model.UserAddress{}
	province := &model.Region{}
	city := &model.Region{}
	district := &model.Region{}
	area := &model.Region{}
	
	err := row.Scan(
		&orderDelivery.Id, &orderDelivery.DeliveryCost, &orderDelivery.DeliveryCostDiscount, &orderDelivery.TotalDeliveryCost, &orderDelivery.TrackingNumber,
		&orderDelivery.Address, &orderDelivery.ReceivedName, &orderDelivery.PhoneNumber, &orderDelivery.Lat, &orderDelivery.Long, &orderDelivery.Status,
		&orderDelivery.DeliveredAt, &orderDelivery.PackageReceivedAt, &courier.Id, &courier.Name, &courier.Category, &courier.Image, &userAddress.PostCode,
		&province.Id, &province.Name, &city.Id, &city.Name, &district.Id, &district.Name, &area.Id, &area.Name, &area.Lat, &area.Long, &area.RefId,
		&orderDelivery.Rate, &orderDelivery.RefId,
	)
	if err != nil {
		return nil, err
	}
	
	area.SetParent(district)
	district.SetParent(city)
	city.SetParent(province)
	
	userAddress.SetProvince(province)
	userAddress.SetCity(city)
	userAddress.SetDistrict(district)
	userAddress.SetArea(area)
	
	orderDelivery.SetCourier(courier)
	orderDelivery.SetUserDelivery(userAddress)
	
	return orderDelivery, nil
}
