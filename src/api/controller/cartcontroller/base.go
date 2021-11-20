package cartcontroller

import (
	"github.com/pandudpn/shopping-cart/src/domain/usecase"
)

type CartController struct {
	CartUsecase usecase.CartUseCaseInterface
}
