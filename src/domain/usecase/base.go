// Package usecase adalah entry point untuk application logic. ini merupakan top level package untuk aplikasi
// top level package hanya menentukan interface, untuk implementasi ditentukan pada sub-package masing-masing
// Ini hanya bergantung dengan package model. tidak boleh package lain selain package `cmd`.
package usecase

import (
	"context"

	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils"
)

// UserUseCaseInterface adalah usecase untuk user melakukan registrasi, login, ataupun yang lainnya
// interface ini support untuk transaction (*sql.Tx)
type UserUseCaseInterface interface {
	// LoginUser adalah case user untuk melakukan login ke dalam aplikasi
	// didalam method ini terdapat berbagai validasi seperti password dan juga email
	LoginUser(ctx context.Context, req *model.RequestLogin) utils.ResponseInterface
	// RegisterUser adalah case user untuk membuat akun baru agar bisa dapat meng-akses aplikasi
	// didalam method ini ada berbagai validasi seperti pengecekan email, no-telp agar tidak terjadi duplicate data
	// method ini juga menggunakan *sql.Tx sebagai query ke database
	// agar jika terjadi error bisa dapat di rollback
	RegisterUser(ctx context.Context, req *model.RequestRegister) utils.ResponseInterface
}

// ProductUseCaseInterface adalah usecase untuk user melakukan pencarian produk atau mengambil data seluruh produk
// ataupun detail produk
type ProductUseCaseInterface interface {
	// GetAllProducts akan mengembalikan seluruh data produk yg masih tersedia dan juga enable
	GetAllProducts(limit, page int, search string) utils.ResponseInterface
	// DetailProductById mengembalikan satu data produk berdasarkan id
	DetailProductById(id int) utils.ResponseInterface
	// DetailProductBySlug mengembalikan satu data produk berdasarkan slug
	DetailProductBySlug(slug string) utils.ResponseInterface
}
