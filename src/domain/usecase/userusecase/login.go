package userusecase

import (
	"context"
	"errors"

	"github.com/pandudpn/shopping-cart/src/api/presenter/userpresenter"
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
	"golang.org/x/crypto/bcrypt"
)

func (uu *UserUseCase) LoginUser(ctx context.Context, req *model.RequestLogin) utils.ResponseInterface {
	user, err := uu.getUserByEmail(req.Email, false)
	if err != nil {
		return userpresenter.ResponseLogin(nil, err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		logger.Log.Error(err)
		err = errors.New("user.password.not_match")

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
