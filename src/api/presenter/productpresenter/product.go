package productpresenter

import (
	"net/http"

	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/repository"
	"github.com/pandudpn/shopping-cart/src/utils"
)

type responseProducts struct {
	Limit        int             `json:"limit"`
	TotalRecords int             `json:"totalRecords"`
	CurrentPage  int             `json:"currentPage"`
	TotalPage    int             `json:"totalPage"`
	Products     []*productsView `json:"products"`
}

type productsView struct {
	Id              int           `json:"id"`
	Name            string        `json:"name"`
	Slug            string        `json:"slug"`
	Description     string        `json:"description"`
	Price           int           `json:"price"`
	DiscountedPrice int           `json:"discountedPrice"`
	Qty             int           `json:"qty"`
	CreatedAt       string        `json:"createdAt"`
	Category        *categoryView `json:"category"`
	Images          []*imageView  `json:"images"`
}

type categoryView struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type imageView struct {
	Url string `json:"url"`
}

var (
	layoutTime        = "2006-01-02T15:04:05-07:00"
	errGlobal  string = "Terjadi kesalahan pada server, silakan coba beberapa saat lagi"

	queryError      string = "query.find.error"
	bodyPayload     string = "body.payload"
	productsSuccess string = "products.success"
	productNotFound string = "product.not_found"

	message = map[string]string{
		queryError:      errGlobal,
		bodyPayload:     "Permintaan kamu tidak lengkap",
		productNotFound: "Produk tidak ditemukan",
	}

	systemCode = map[string]string{
		productsSuccess: "30",
		productNotFound: "32",

		bodyPayload: "80",
		queryError:  "81",
	}

	statusCode = map[string]int{
		productsSuccess: http.StatusOK,
		productNotFound: http.StatusNotFound,
		bodyPayload:     http.StatusBadRequest,
		queryError:      http.StatusInternalServerError,
	}
)

func ResponseProducts(value interface{}, err error, redis repository.RedisRepositoryInterface) utils.ResponseInterface {
	if err != nil {
		errString := err.Error()
		return utils.Error(statusCode[errString], systemCode[errString], message[errString], err)
	}

	if key, ok := value.(map[string]interface{}); ok {
		res := responseProducts{
			Limit:        key["limit"].(int),
			TotalRecords: key["totalRecord"].(int),
			CurrentPage:  key["currentPage"].(int),
			TotalPage:    int(key["totalPage"].(float64)),
		}
		if products, ok := key["products"].([]*model.Product); ok {
			productViews := createProductsView(products)
			res.Products = productViews
		}

		// go redis.SaveProductsCache(res, key["searchProduct"].(string))
		return utils.Success(statusCode[productsSuccess], systemCode[productsSuccess], res)
	}

	product := value.(*model.Product)
	pv := productView(product)

	return utils.Success(statusCode[productsSuccess], systemCode[productsSuccess], pv)
}

func createProductsView(products []*model.Product) []*productsView {
	var productViews = make([]*productsView, 0)

	for _, product := range products {
		pv := productView(product)
		productViews = append(productViews, pv)
	}

	return productViews
}

func productView(product *model.Product) *productsView {
	productView := &productsView{
		Id:              product.Id,
		Name:            product.Name,
		Slug:            product.Slug,
		Description:     *product.Description,
		Price:           int(product.Price),
		DiscountedPrice: int(product.DiscountedPrice),
		Qty:             product.GetQuantity(),
		CreatedAt:       product.CreatedAt.Format(layoutTime),
	}

	if product.Category != nil {
		category := &categoryView{
			Id:   product.Category.Id,
			Name: product.Category.Name,
			Slug: product.Category.Slug,
		}
		productView.Category = category
	}

	if len(product.GetImages()) > 0 {
		images := make([]*imageView, 0)
		for _, image := range product.GetImages() {
			img := &imageView{
				Url: image.GetImage().GetFile(),
			}

			images = append(images, img)
		}

		productView.Images = images
	}

	return productView
}
