package postgre

import (
	"github.com/private-project-pp/product-rpc-service/domain"
	"github.com/private-project-pp/product-rpc-service/entity"
	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func SetupProductsRepo(db *gorm.DB) domain.Product {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) AddProduct(in entity.Product) (err error) {
	return nil
}

func (r *productRepo) BulkAddProductImages(in []entity.ProductsImages) (err error) {
	return nil
}
