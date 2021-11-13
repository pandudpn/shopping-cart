package productusecase

import (
	"github.com/pandudpn/shopping-cart/src/repository"
)

type ProductUseCase struct {
	Redis        repository.RedisRepositoryInterface
	ProductRepo  repository.ProductRepositoryInterface
	ImageRepo    repository.ProductImageRepositoryInterface
	CategoryRepo repository.ProductCategoryRepositoryInterface
}
