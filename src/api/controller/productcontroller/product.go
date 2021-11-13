package productcontroller

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (pc *ProductController) GetProductsHandler(e echo.Context) error {
	search := e.QueryParam("search")
	limit, _ := strconv.Atoi(e.QueryParam("limit"))
	page, _ := strconv.Atoi(e.QueryParam("page"))

	return pc.ProductUseCase.GetAllProducts(limit, page, search).JSON(e)
}

func (pc *ProductController) DetailProductHandler(e echo.Context) error {
	id := e.Param("id")

	productId, err := strconv.Atoi(id)
	if err != nil {
		logger.Log.Debugf("url param bukan integer %s", id)
		return pc.ProductUseCase.DetailProductBySlug(id).JSON(e)
	}

	return pc.ProductUseCase.DetailProductById(productId).JSON(e)
}
