package postgre

import (
	"github.com/private-project-pp/product-rpc-service/domain"
	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

func SetupProductsRepo(db *gorm.DB) domain.Product {
	return &users{
		db: db,
	}
}
