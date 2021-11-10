package userusecase

import "github.com/pandudpn/shopping-cart/src/repository"

type UserUseCase struct {
	UserRepo  repository.UserRepositoryInterface
	TxRepo    repository.TxRepositoryInterface
	RedisRepo repository.RedisRepositoryInterface
}
