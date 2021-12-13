package rdb

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
	"github.com/spf13/viper"
)

func (r *RedisRepository) SaveProductsCache(i interface{}, searchProduct string) error {
	var (
		key  string
		data = make(map[string]interface{})
	)

	iByte, err := json.Marshal(i)
	if err != nil {
		logger.Log.Debug(err)
		return err
	}

	err = json.Unmarshal(iByte, &data)
	if err != nil {
		logger.Log.Debug(err)
		return err
	}

	duration, err := time.ParseDuration(viper.GetString("cached.products"))
	if err != nil {
		logger.Log.Error(err)
		return err
	}

	if limit, found := data["limit"].(float64); found {
		key = fmt.Sprintf("product-%.0f", limit)
	}

	if page, found := data["currentPage"].(float64); found {
		key = fmt.Sprintf("%s-%.0f", key, page)
	}

	// checking, if not null, then insert into key
	if !reflect.ValueOf(searchProduct).IsZero() {
		key = fmt.Sprintf("%s-%s", key, searchProduct)
		// delete data searchProduct from map
	}

	err = r.save(key, data, duration)
	if err != nil {
		logger.Log.Error(err)
		return err
	}

	return nil
}

func (r *RedisRepository) SetCachedShipper(cart *model.Cart, dataByte []byte) error {
	var i interface{}

	keyString := fmt.Sprintf("%d%d", cart.Id, *cart.UserAddressId)
	key := base64.StdEncoding.EncodeToString([]byte(keyString))

	duration, err := time.ParseDuration(viper.GetString("cached.shipper"))
	if err != nil {
		logger.Log.Errorf("error parsing duration shipper.id %v", err)
		return err
	}

	err = json.Unmarshal(dataByte, &i)
	if err != nil {
		logger.Log.Errorf("error unmarshal on cached shipper %v", err)
		return err
	}

	err = r.save(key, i, duration)
	if err != nil {
		logger.Log.Errorf("error save to redis %v", err)
		return err
	}

	return nil
}

func (r *RedisRepository) GetCachedShipper(cart *model.Cart) ([]byte, error) {
	keyString := fmt.Sprintf("%d%d", cart.Id, *cart.UserAddressId)

	key := base64.StdEncoding.EncodeToString([]byte(keyString))

	result, err := r.getString(key)
	if err != nil {
		return nil, err
	}

	return []byte(result), nil
}
