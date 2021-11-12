package repository

import "github.com/pandudpn/shopping-cart/src/domain/model"

type UserRepositoryInterface interface {
	// FindById mencari user berdasarkan id
	FindById(id int) (*model.User, error)
	// FindByEmail mencari user berdasarkan email
	// ini bisa diigunakan untuk login ataupun register
	FindByEmail(email string) (*model.User, error)
	// FindByPhone mencari user bedasarkan nomer telpon
	// digunakan untuk register agar tidak terjadi duplicate data
	FindByPhone(phone string) (*model.User, error)
	// InsertUser akan membuat baris baru pada table user
	InsertUser(user *model.User) error
}

// RedisRepositoryInterface digunakan untuk kumpulan query-query yang langsung ke redis db
type RedisRepositoryInterface interface {
	// SetSession digunakan ketika user ingin melakukan login
	// menyimpan data user seperti `id`, `email` dan juga `name`
	SetSession(user *model.User) (string, error)
	// GetSession untuk mengambil data session dari redis
	// biasanya method ini digunakan pada level middleware
	GetSession(key string) (*model.User, error)
}

// ProductRepositoryInterface digunakan untuk kumpulan query untuk table product
type ProductRepositoryInterface interface {
	// FindAllProducts akan menampilkan seluruh data product
	// method ini belum support dynamic sorting dan juga pagination
	// jika ingin menggunakan dynamic sorting ataupun pagination, silahkan tambahkan sendiri
	FindAllProducts() ([]*model.Product, error)
	// FindProductById akan menampilkan single data product berdasarkan id
	FindProductById(id int) (*model.Product, error)
	// FindProductBySlug akan menampilkan single data product berdasarkan slug
	FindProductBySlug(slug string) (*model.Product, error)
	// FindProductsByName akan data-data produk berdasarkan produk yg dicari (search product)
	FindProductsByName(name string) ([]*model.Product, error)
}

// ProductImageRepositoryInterface adalah kumpulan query-query untuk mengambil image product
type ProductImageRepositoryInterface interface {
	// FindImagesByProductId akan mengembalikan array data product image yang sudah di relasikan
	// ke table media_file
	FindImagesByProductId(productId int) ([]*model.ProductImage, error)
}

type ProductCategoryRepositoryInterface interface{}

// TxRepositoryInterface untuk melakukan transactional database dengan interface2 lainnya
type TxRepositoryInterface interface {
	TxEnd(txFunc func() error) error
}
