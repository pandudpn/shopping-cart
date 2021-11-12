package productusecase

import (
	"github.com/pandudpn/shopping-cart/src/repository"
)

type ProductUseCase struct {
	ProductRepo  repository.ProductRepositoryInterface
	CategoryRepo repository.ProductCategoryRepositoryInterface
}
