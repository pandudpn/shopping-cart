package userusecase

import (
	"context"

	"github.com/pandudpn/shopping-cart/src/domain/model"
)

func (uu *UserUseCase) LoginUser(ctx context.Context, user *model.User) (*model.User, error) {
	user, err := uu.UserRepo.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
