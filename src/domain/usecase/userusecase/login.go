package userusecase

import (
	"context"

	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (uu *UserUseCase) LoginUser(ctx context.Context, user *model.User) (*model.User, error) {
	user, err := uu.UserRepo.FindByEmail(user.Email)
	if err != nil {
		logger.Log.Errorf("error get user by email %v", err)
		return nil, err
	}

	return user, nil
}
