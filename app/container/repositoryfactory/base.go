// Package repositoryfactory menggunakan factory method pattern untuk membuat tipe yang kongkret.
package repositoryfactory

import (
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
)

// untuk me-mapping "repository code"
// masing-masing repositoryFactory mempunyai code dan juga builder/factory sendiri
var rFbMap = map[string]repositoryFbInterface{
	constant.USER:                &userRepositoryFactory{},
	constant.REDIS:               &redisRepositoryFactory{},
	constant.TX:                  &txRepositoryFactory{},
	constant.PRODUCT:             &productRepositoryFactory{},
	constant.PRODUCTCATEGORY:     &productCategoryRepositoryFactory{},
	constant.PRODUCTIMAGE:        &productImageRepositoryFactory{},
	constant.MEDIAFILE:           &mediaFileRepositoryFactory{},
	constant.CART:                &cartRepositoryFactory{},
	constant.CART_PRODUCT:        &cartProductRepositoryFactory{},
	constant.STOCK:               &stockRepositoryFactory{},
	constant.COURIER:             &courierRepositoryFactory{},
	constant.PAYMENTMETHOD:       &paymentMethodRepositoryFactory{},
	constant.USERADDRESS:         &userAddressRepositoryFactory{},
	constant.ORDER:               &orderRepositoryFactory{},
	constant.ORDERPRODUCT:        &orderProductRepositoryFactory{},
	constant.ORDERPAYMENT:        &orderPaymentRepositoryFactory{},
	constant.ORDERDELIVERY:       &orderDeliveryRepositoryFactory{},
	constant.ORDERDELIVERYSTATUS: &orderDeliveryStatusRepositoryFactory{},
}

type RepositoryFactoryInterface interface{}

// Builder interface untuk factory repository
// setiap factory yg dibuat harus mengimplementasikan method ini
type repositoryFbInterface interface {
	Build(c container.Container) (RepositoryFactoryInterface, error)
}

// GetRepositoryFbMap adalah aksesor untuk mengambil factory builder
func GetRepositoryFbMap(code string) repositoryFbInterface {
	return rFbMap[code]
}
