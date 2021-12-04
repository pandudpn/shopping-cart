package useraddressrepo

import "github.com/pandudpn/shopping-cart/src/domain/model"

func (uar *UserAddressRepository) FindDefaultDeliveryByUser(user *model.User) (*model.UserAddress, error) {
	row := uar.DB.QueryRow(QUERY_BY_DELIVERY_DEFAULT, user.Id)

	userAddress, err := rowToUserAddress(row)
	if err != nil {
		return nil, err
	}

	userAddress.SetUser(user)
	return userAddress, nil
}

func (uar *UserAddressRepository) FindAllByUser(user *model.User) ([]*model.UserAddress, error) {
	userAddresses := make([]*model.UserAddress, 0)

	rows, err := uar.DB.Query(QUERY_BY_USER, user.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		userAddress, err := rowsToUserAddress(rows)
		if err != nil {
			return nil, err
		}

		userAddress.SetUser(user)
		userAddresses = append(userAddresses, userAddress)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return userAddresses, nil
}

func (uar *UserAddressRepository) FindUserAddressById(id int) (*model.UserAddress, error) {
	row := uar.DB.QueryRow(QUERY_BY_ID, id)

	userAddress, err := rowToUserAddress(row)
	return userAddress, err
}
