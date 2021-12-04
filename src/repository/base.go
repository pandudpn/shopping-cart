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

// UserAddressRepositoryInterface adalah kumpulan query-query untuk table 'user_address'
type UserAddressRepositoryInterface interface {
	// FindAllByUser digunakan untuk mengambil seluruh list user_addresss berdasarkan
	// user.id
	FindAllByUser(user *model.User) ([]*model.UserAddress, error)
	// FindDefaultDeliveryByUser akan mendapatkan satu data delivery address
	// yg sudah menjadi default
	// method ini nantinya digunakan untuk membuat default cart.address_id ketika
	// pertama kali melakukan checkout
	FindDefaultDeliveryByUser(user *model.User) (*model.UserAddress, error)
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
	// FindCartProductByCartIdAndProductId digunakan untuk pengecekan, apakah
	// produk id dan juga cart id tersebut sudah ada atau belum
	// jika belum ada, maka akan menjalankan method `InsertNewCartProduct`
	// jika sudah ada, maka akan menjalankan method `UpdateCartProduct`
	FindCartProductByCartIdAndProductId(cartId, productId int) (*model.CartProduct, error)
	// InsertNewCartProduct menambahkan baris baru pada 'cart_products' dan mengembalikan
	// id yg baru saja dibuat
	InsertNewCartProduct(cartProduct *model.CartProduct) error
	// UpdateCartProduct digunakan untuk merubah data-data seperti qty, dan total_price
	// berdasarkan id dari cart_product itu sendiri
	UpdateCartProduct(cartProduct *model.CartProduct) error
}

// StockRepositoryInterface adalah kumpulan query-query untuk mengambil data stock
type StockRepositoryInterface interface {
	// FindStockByProductId akan mengambil satu data stock berdasarkan product_id
	FindStockByProductId(product *model.Product) (*model.Stock, error)
	// UpdateStock akan melakukan update stock terbaru
	UpdateStock(stockId, qty int) error
}

// PaymentMethodRepositoryInterface adalah kumpulan query-query pada table 'payment_method'
type PaymentMethodRepositoryInterface interface {
	// FindEnabledPaymentMethod akan mengembalikan seluruh data metode pembayaran yang aktif
	FindEnabledPaymentMethod() ([]*model.PaymentMethod, error)
	// FindPaymentMethodById digunakan untuk mengambil satu data berdasarkan payment_method.id
	FindPaymentMethodById(paymentMethodId int) (*model.PaymentMethod, error)
}

// CourierRepositoryInterface adalah kumpulan query-query pada table 'courier'
type CourierRepositoryInterface interface {
	// FindenabledCourier untuk mendapatkan seluruh kurir yg aktif
	FindEnabledCourier() ([]*model.Courier, error)
	// FindCourierById digunakan untuk mendapatkan single data kurir
	// berdasarkan courier.id
	FindCourierById(courierId int) (*model.Courier, error)
}

// TxRepositoryInterface untuk melakukan transactional database dengan interface2 lainnya
type TxRepositoryInterface interface {
	TxEnd(txFunc func() error) error
}
