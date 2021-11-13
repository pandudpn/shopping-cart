package rdb

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

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
