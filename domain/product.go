package domain

import (
	"github.com/private-project-pp/product-rpc-service/entity"
	"gorm.io/gorm"
)

type Product interface {
	AddProduct(tx *gorm.DB, in entity.Product) (err error)
	BulkAddProductImages(tx *gorm.DB, in []entity.ProductsImages) (err error)
}
