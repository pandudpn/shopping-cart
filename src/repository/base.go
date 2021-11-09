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
