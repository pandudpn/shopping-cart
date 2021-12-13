// Package rdb adalah sub-package dari repository package.
// di package ini khusus digunakan untuk Redis query
// e.g: UserSession, OTP, ProductList, Detail, e.t.c
package rdb

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

type RedisRepository struct {
	RDb dbc.RDbc
}

func (rr *RedisRepository) save(key string, value interface{}, duration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	dByte, err := json.Marshal(value)
	if err != nil {
		logger.Log.Errorf("error save to redis %v", err)
		return err
	}

	err = rr.RDb.Set(ctx, key, string(dByte), duration).Err()
	return err
}

func (rr *RedisRepository) get(key string, dest interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	res := rr.RDb.Get(ctx, key)
	if res.Err() != nil {
		return res.Err()
	}

	err := res.Scan(dest)
	return err
}

func (rr *RedisRepository) getString(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	res := rr.RDb.Get(ctx, key)
	if res.Err() != nil {
		return "", res.Err()
	}

	return res.Val(), nil
}

func (rr *RedisRepository) delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := rr.RDb.Del(ctx, key).Err()
	return err
}

func (rr *RedisRepository) saveOneDay(key string, value interface{}) error {
	timezone, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		logger.Log.Error(err)
		return err
	}

	now := time.Now().In(timezone)
	nextDay := now.AddDate(0, 0, 1)
	nextDay = time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, timezone)

	duration := nextDay.Sub(now).Seconds()

	return rr.save(key, value, time.Duration(duration)*time.Second)
}
