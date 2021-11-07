// Package usecasefactory menggunakan design pattern factory method untuk membuat berbagai macam bisnis logic
// masing-masing usecase membutuhkan factory sendiri
package usecasefactory

import (
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
)

// untuk me-mapping "usecase code" dari configuration yaml file
// masing-masing factory mempunyai package/file itu sendiri
// contoh : `user` akan mempunyai 2 case yaitu `register` dan `login`
// maka `user` tersebut akan mempunyai package bernama `userusecasefactory`
var ucFbMap = map[string]useCaseFbInterface{
	constant.USER: &userUseCaseFactory{},
}

type UseCaseFactoryInterface interface{}

// Builder interface untuk factory usecase
// setiap factory yg dibuat harus mengimplementasikan method ini
type useCaseFbInterface interface {
	Build(c container.Container) (UseCaseFactoryInterface, error)
}

func GetUseCaseFbMap(code string) useCaseFbInterface {
	return ucFbMap[code]
}
