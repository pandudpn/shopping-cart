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
	// SaveProductsCache digunakan untuk menyimpan data product berdasarkan limit dan juga pagination
	// hasil dari key tersebut berupa `products-{limit}-{page}`
	// jika terdapat pencarian product maka key menjadi `products-{limit}-{page}-{search}`
	SaveProductsCache(i interface{}, searchProduct string) error
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

// CartRepositoryInterface adalah kumpulan query-query untuk mengambil/menambahkan/mengubah data cart
type CartRepositoryInterface interface {
	// FindActiveCartByUserId mengembalikan cart yg masih aktif pada user tersebut
	// (cart teerakhir yg aktif)
	FindActiveCartByUserId(userId int) (*model.Cart, error)
	// FindCartByKey digunakan untuk mengambil data cart untuk melakukan checkout data
	FindCartByKey(key string) (*model.Cart, error)
	// InsertNewCart untuk membuat keranjang belanja baru jika tidak ada yg sedang aktif
	InsertNewCart(cart *model.Cart) error
}

// CartProductRepositoryInterface adalah kumpulan query-query pada table 'cart_product'
type CartProductRepositoryInterface interface {
	// FindCartProductsByCartId akan mengambil data 'cart_products' berdasarkan cart_id nya
	// dan akan langsung di inject ke cart secara langsung
	FindCartProductsByCartId(cart *model.Cart) error
	// InsertNewCartProduct menambahkan baris baru pada 'cart_products' dan mengembalikan
	// id yg baru saja dibuat
	InsertNewCartProduct(cartProduct *model.CartProduct) error
}

// StockRepositoryInterface adalah kumpulan query-query untuk mengambil data stock
type StockRepositoryInterface interface {
	// FindStockByProductId akan mengambil satu data stock berdasarkan product_id
	FindStockByProductId(product *model.Product) (*model.Stock, error)
	// UpdateStock akan melakukan update stock terbaru
	UpdateStock(stockId, qty int) error
}

// TxRepositoryInterface untuk melakukan transactional database dengan interface2 lainnya
type TxRepositoryInterface interface {
	TxEnd(txFunc func() error) error
}
