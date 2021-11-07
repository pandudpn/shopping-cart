// Package datastorefactory menggunakan pola metode factory untuk membuat database handler
// Datastore menyediakan akses data untuk package `model`
// hanya menggunakan satu method `Build()` untuk membuat berbagai tipe datastore
package datastorefactory

import (
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
)

// untuk me-mapping database code
var dsFbMap = map[string]dataStoreFbInterface{
	constant.PSQL: &psqlFactory{},
}

type DataStoreInterface interface{}

// Builder interface untuk factory database handler
// setiap factory yg dibuat harus mengimplementasikan method ini
type dataStoreFbInterface interface {
	Build(c container.Container, enableTx bool) (DataStoreInterface, error)
}

// GetDataStoreFbMap adalah aksesor untuk mengambil factory builder
func GetDataStoreFbMap(code string) dataStoreFbInterface {
	return dsFbMap[code]
}
