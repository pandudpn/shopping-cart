package productusecase

import (
	"errors"
	"math"
	"reflect"

	"github.com/pandudpn/shopping-cart/src/api/presenter/productpresenter"
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (puu *ProductUseCase) GetAllProducts(limit, page int, search string) utils.ResponseInterface {
	var (
		products = make([]*model.Product, 0)
		res      = make(map[string]interface{})
		offset   int
	)

	if reflect.ValueOf(limit).IsZero() {
		limit = 20
	}
	res["limit"] = limit

	if reflect.ValueOf(page).IsZero() {
		page = 1
	}
	res["currentPage"] = page

	if reflect.ValueOf(search).IsZero() {
		p, err := puu.ProductRepo.FindAllProducts()
		if err != nil {
			logger.Log.Errorf("error get all products %v", err)
			err = errors.New("query.find.error")
			return productpresenter.ResponseProducts(nil, err, nil)
		}

		for _, product := range p {
			images, err := puu.ImageRepo.FindImagesByProductId(product.Id)
			if err != nil {
				logger.Log.Error(err)
				continue
			}

			product.SetImages(images)
		}

		products = p
	} else {
		p, err := puu.ProductRepo.FindProductsByName(search)
		if err != nil {
			logger.Log.Errorf("error get products by search %v", err)
			err = errors.New("query.find.error")
			return productpresenter.ResponseProducts(nil, err, nil)
		}

		for _, product := range p {
			images, err := puu.ImageRepo.FindImagesByProductId(product.Id)
			if err != nil {
				logger.Log.Error(err)
				continue
			}

			product.SetImages(images)
		}

		products = p
	}
	res["totalRecord"] = len(products)

	totalPage := math.Ceil(float64(len(products)) / float64(limit))
	res["totalPage"] = totalPage

	offset = (page - 1) * limit
	limit = limit * page

	if limit > len(products) {
		limit = len(products)
	}

	if offset > len(products) {
		offset = len(products)
	}

	products = products[offset:limit]

	res["searchProduct"] = search
	res["products"] = products

	return productpresenter.ResponseProducts(res, nil, puu.Redis)
}

func (puu *ProductUseCase) DetailProductById(id int) utils.ResponseInterface {
	product, err := puu.ProductRepo.FindProductById(id)
	if err != nil {
		logger.Log.Error(err)
		err = errors.New("product.not_found")
		return productpresenter.ResponseProducts(nil, err, nil)
	}

	return puu.returnDetailProduct(product)
}

func (puu *ProductUseCase) DetailProductBySlug(slug string) utils.ResponseInterface {
	product, err := puu.ProductRepo.FindProductBySlug(slug)
	if err != nil {
		logger.Log.Error(err)
		err = errors.New("product.not_found")
		return productpresenter.ResponseProducts(nil, err, nil)
	}

	return puu.returnDetailProduct(product)
}

func (puu *ProductUseCase) returnDetailProduct(product *model.Product) utils.ResponseInterface {
	images, err := puu.ImageRepo.FindImagesByProductId(product.Id)
	if err != nil {
		logger.Log.Error(err)
	} else {
		product.SetImages(images)
	}

	return productpresenter.ResponseProducts(product, nil, nil)
}
