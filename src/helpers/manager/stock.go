package manager

import (
	"github.com/pandudpn/shopping-cart/src/repository"
)

type stockManager struct {
	productRepo repository.ProductRepositoryInterface
	stockRepo   repository.StockRepositoryInterface
}

func NewStockManager(pr repository.ProductRepositoryInterface, sr repository.StockRepositoryInterface) StockManagerInterface {
	return &stockManager{
		productRepo: pr,
		stockRepo:   sr,
	}
}
