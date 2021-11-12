package productcontroller

import (
	"github.com/pandudpn/shopping-cart/src/domain/usecase"
)

type ProductController struct {
	ProductUseCase usecase.ProductUseCaseInterface
}
