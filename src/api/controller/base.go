// Package controller adalah top level dari package controller/handler dari routing
// package untuk membuat response sukses atau gagal berdasarkan bisnis logic yg dibutuhkan
//
// package ini digunakan untuk menerima data masuk dan melakukan parsing data
// sesuai dengan kebutuhan dari masing-masing bisnis logic nya.
// jika data yang dibutuhkan tidak sesuai, akan mengembalikan error berupa response error
package controller

import "github.com/labstack/echo"

// UserControllerInterface adalah kumpulan controller/handler user
type UserControllerInterface interface {
	LoginHandler(e echo.Context) error
	RegisterHandler(e echo.Context) error
}

// ProductControllerInterface adalah kumpulan controller/handler products
type ProductControllerInterface interface {
	GetProductsHandler(e echo.Context) error
}
