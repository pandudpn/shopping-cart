package userusecase

import "github.com/pandudpn/shopping-cart/src/repository"

type UserUseCase struct {
	UserRepo  repository.UserRepositoryInterface
	RedisRepo repository.RedisRepositoryInterface
}
