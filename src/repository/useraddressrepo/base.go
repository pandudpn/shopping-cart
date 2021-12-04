package useraddressrepo

import (
	"database/sql"

	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

const (
	QUERY_SELECT_ENABLED = "select ua.id, ua.user_id, ua.receiver_name, ua.phone_number, ua.post_code, " +
		`ua.type, ua.address, ua."default", ua.lat, ua.long, p.id, p.parent_id, p.name, ` +
		"p.ref_id, p.lat, p.long, c.id, c.parent_id, c.ref_id, c.name, c.lat, c.long, " +
		"d.id, d.parent_id, d.ref_id, d.name, d.lat, d.long, a.id, " +
		"a.parent_id, a.ref_id, a.name, a.lat, a.long " +
		"from public.user_address ua " +
		"inner join public.region p on ua.province_id = p.id " +
		"inner join public.region c on ua.city_id  = c.id " +
		"inner join public.region d on ua.district_id = d.id " +
		"inner join public.region a on ua.area_id = a.id " +
		"where ua.deleted = false "
	QUERY_BY_ID               = QUERY_SELECT_ENABLED + "and ua.id=$1"
	QUERY_BY_USER             = QUERY_SELECT_ENABLED + "and ua.user_id = $1"
	QUERY_BY_DELIVERY_DEFAULT = QUERY_BY_USER + ` and ua."default" = true and type = 'delivery'`
)

type UserAddressRepository struct {
	DB dbc.SqlDbc
}

func rowsToUserAddress(rows *sql.Rows) (*model.UserAddress, error) {
	userAddress := &model.UserAddress{}
	province := &model.Region{}
	city := &model.Region{}
	district := &model.Region{}
	area := &model.Region{}

	err := rows.Scan(
		&userAddress.Id, &userAddress.UserId, &userAddress.ReceiverName, &userAddress.PhoneNumber,
		&userAddress.PostCode, &userAddress.Type, &userAddress.Address, &userAddress.Default,
		&userAddress.Lat, &userAddress.Long, &province.Id, &province.ParentId, &province.Name,
		&province.RefId, &province.Lat, &province.Long, &city.Id, &city.ParentId, &city.RefId,
		&city.Name, &city.Lat, &city.Long, &district.Id, &district.ParentId, &district.RefId,
		&district.Name, &district.Lat, &district.Long, &area.Id, &area.ParentId, &area.RefId,
		&area.Name, &area.Lat, &area.Long,
	)
	if err != nil {
		return nil, err
	}

	city.SetParent(district)
	district.SetParent(city)
	city.SetParent(province)

	userAddress.SetProvince(province)
	userAddress.SetCity(city)
	userAddress.SetDistrict(district)
	userAddress.SetArea(area)

	return userAddress, nil
}

func rowToUserAddress(row *sql.Row) (*model.UserAddress, error) {
	userAddress := &model.UserAddress{}
	province := &model.Region{}
	city := &model.Region{}
	district := &model.Region{}
	area := &model.Region{}

	err := row.Scan(
		&userAddress.Id, &userAddress.UserId, &userAddress.ReceiverName, &userAddress.PhoneNumber,
		&userAddress.PostCode, &userAddress.Type, &userAddress.Address, &userAddress.Default,
		&userAddress.Lat, &userAddress.Long, &province.Id, &province.ParentId, &province.Name,
		&province.RefId, &province.Lat, &province.Long, &city.Id, &city.ParentId, &city.RefId,
		&city.Name, &city.Lat, &city.Long, &district.Id, &district.ParentId, &district.RefId,
		&district.Name, &district.Lat, &district.Long, &area.Id, &area.ParentId, &area.RefId,
		&area.Name, &area.Lat, &area.Long,
	)
	if err != nil {
		return nil, err
	}

	city.SetParent(district)
	district.SetParent(city)
	city.SetParent(province)

	userAddress.SetProvince(province)
	userAddress.SetCity(city)
	userAddress.SetDistrict(district)
	userAddress.SetArea(area)

	return userAddress, nil
}
