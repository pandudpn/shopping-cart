package userusecase

import (
	"context"

	"github.com/pandudpn/shopping-cart/src/api/presenter/userpresenter"
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils"
)

func (uu *UserUseCase) RegisterUser(ctx context.Context, req *model.RequestRegister) utils.ResponseInterface {
	err := uu.TxRepo.TxEnd(func() error {
		return uu.register(req)
	})

	if err != nil {
		return userpresenter.ResponseRegister(nil, err)
	}

	return userpresenter.ResponseRegister("register.success", nil)
}
