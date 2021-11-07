// Package middleware adalah top level package dari middleware
// top level package hanya menentukan interface. untuk implementasi dilakukan pada sub-packaege masing-masing
package middleware

import "net/http"

// CachedMiddlewareInterface adalah sub-package dari package middleware
// digunakan untuk meng-handle data-data cache seperti session login, otp, dll
type CachedMiddlewareInterface interface {
	CheckSession(next http.Handler) http.Handler
}
