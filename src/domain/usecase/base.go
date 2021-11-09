// Package usecase adalah entry point untuk application logic. ini merupakan top level package untuk aplikasi
// top level package hanya menentukan interface, untuk implementasi ditentukan pada sub-package masing-masing
// Ini hanya bergantung dengan package model. tidak boleh package lain selain package `cmd`.
package usecase

import (
	"context"

	"github.com/pandudpn/shopping-cart/src/utils"
)

// UserUseCaseInterface adalah usecase untuk user melakukan registrasi, login, ataupun yang lainnya
// interface ini support untuk transaction (*sql.Tx)
type UserUseCaseInterface interface {
	LoginUser(ctx context.Context, email string) utils.ResponseInterface
}
