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

// CartUseCaseInterface adalah usecase untuk user menambahkan sebuah produk atau mengambil data-data produk
// yang ada pada keranjang mereka.
// usecase ini support untuk *sql.Tx
type CartUseCaseInterface interface {
	// AddToCart adalah sebuah method untuk user melakukan penambahan sebuah produk ke keranjang user tersebut
	AddToCart(ctx context.Context, req *model.RequestAddToCart) utils.ResponseInterface
	// GetCart adalah sebuah method untuk user menampilkan produk-produk yg telah dimasukkan kedalam ke keranjang belanja
	GetCart(ctx context.Context, userId int, key string) utils.ResponseInterface
}

// CheckoutUseCaseInterface adalah usecase untuk user melakukan pemilihan kurir, metode pembayaran serta alamat yang dikirim
// pada usecase ini terdapat banyak beberapa logic seperti
//    1. mencari kurir berdasarkan alamat yg dituju
//    2. mencari metode pembayaran yg bisa dilakukan
//    3. apakah barang tersebut available dibeli atau tidak
type CheckoutUseCaseInterface interface {
	// GetCheckout adalah bisnis logic untuk user pertama kali masuk ke page halaman pembayaran atau checkout
	// pada method ini, data akan direset semua untuk meng-clear kan data yg akan diupdate oleh user itu sendiri seperti
	// Courier, PaymentMethod
	GetCheckout(ctx context.Context, key string, userId int) utils.ResponseInterface
	// Update adalah bisnis logic untuk user memilih (mengubah data) kurir yg ingin dikirim, alamat tujuan, serta metode pembayaran
	// yg user inginkan.
	// method ini digunakan untuk mengubah data sesuai yg di user inginkan
	Update(ctx context.Context, req *model.RequestCheckout) utils.ResponseInterface
}
