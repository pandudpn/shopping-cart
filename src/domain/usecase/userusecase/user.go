package userusecase

import (
	"database/sql"
	"errors"

	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (uu *UserUseCase) getUserByEmail(email string, isRegister bool) (*model.User, error) {
	user, err := uu.UserRepo.FindByEmail(email)
	if err != nil {
		logger.Log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			if isRegister {
				return model.NewUser(), nil
			}
			return nil, errors.New("user.email.not_found")
		}

		return nil, errors.New("query.find.error")
	}

	if !user.IsVerified() || !user.Enabled {
		return nil, errors.New("user.not_active")
	}

	return user, nil
}

func (uu *UserUseCase) getUserByPhone(phone string, isRegister bool) (*model.User, error) {
	user, err := uu.UserRepo.FindByPhone(phone)
	if err != nil {
		logger.Log.Error(err)

		if errors.Is(err, sql.ErrNoRows) {
			if isRegister {
				return model.NewUser(), nil
			}
			return nil, errors.New("user.phone.not_found")
		}

		return nil, errors.New("query.find.error")
	}

	if !user.IsVerified() || !user.Enabled {
		return nil, errors.New("user.not_active")
	}

	return user, nil
}
