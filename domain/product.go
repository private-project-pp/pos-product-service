package domain

import "github.com/private-project-pp/product-rpc-service/entity"

type Product interface {
	AddProduct(in entity.Product) (err error)
	BulkAddProductImages(in []entity.ProductsImages) (err error)
}
