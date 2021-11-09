package rdb

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
	"github.com/spf13/viper"
)

func (rr *RedisRepository) SetSession(user *model.User) (string, error) {
	expired, err := time.ParseDuration(viper.GetString("session.expired"))
	if err != nil {
		return "", err
	}

	key := uuid.New().String()
	key = strings.ReplaceAll(key, "-", "")

	data := map[string]interface{}{
		"id":    user.Id,
		"email": user.Email,
		"name":  user.Name,
	}

	err = rr.save(key, data, expired)
	return key, err
}

func (rr *RedisRepository) GetSession(key string) (*model.User, error) {
	user := model.User{}

	err := rr.get(key, &user)
	if err != nil {
		logger.Log.Errorf("data not found in redis %v", err)
		return nil, err
	}

	return &user, nil
}
