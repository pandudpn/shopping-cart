package userusecase

import (
	"context"
	"errors"

	"github.com/pandudpn/shopping-cart/src/api/presenter/userpresenter"
	"github.com/pandudpn/shopping-cart/src/utils"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (uu *UserUseCase) LoginUser(ctx context.Context, email string) utils.ResponseInterface {
	user, err := uu.getUserByEmail(email, false)
	if err != nil {
		return userpresenter.ResponseLogin(nil, err)
	}

	token, err := uu.RedisRepo.SetSession(user)
	if err != nil {
		logger.Log.Errorf("error create session %v", err)
		err = errors.New("session.create.error")

		return userpresenter.ResponseLogin(nil, err)
	}

	res := map[string]interface{}{
		"tokenType":   "Bearer",
		"accessToken": token,
	}

	return userpresenter.ResponseLogin(res, nil)
}
