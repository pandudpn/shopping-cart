package productcontroller

import (
	"strconv"

	"github.com/labstack/echo"
)

func (pc *ProductController) GetProductsHandler(e echo.Context) error {
	search := e.QueryParam("search")
	limit, _ := strconv.Atoi(e.QueryParam("limit"))
	page, _ := strconv.Atoi(e.QueryParam("page"))

	return pc.ProductUseCase.GetAllProducts(limit, page, search).JSON(e)
}
