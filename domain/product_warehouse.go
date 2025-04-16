package domain

import "github.com/private-project-pp/product-rpc-service/entity"

type ProductWarehouse interface {
	GetProductByID(id string) (out entity.ProductsWarehouse, err error)
	AddProduct(in entity.ProductsWarehouse) (err error)
	UpdateProduct(in entity.ProductsWarehouse) (err error)
}
