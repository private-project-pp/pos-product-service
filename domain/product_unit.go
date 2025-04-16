package domain

import "github.com/private-project-pp/product-rpc-service/entity"

type ProductUnit interface {
	GetProductUnitById(in string) (out entity.UnitOfMeasuremnet, err error)
	AddProductUnit(in entity.UnitOfMeasuremnet) (err error)
	UpdateProductUnit(in entity.UnitOfMeasuremnet) (err error)
	GetProductUnitByProductAndUnitId(productId, unitId, olderUnitId string) (out entity.ProductsUnitComparator, err error)
}
