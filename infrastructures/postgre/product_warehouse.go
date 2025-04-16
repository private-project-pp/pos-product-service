package postgre

import (
	"github.com/private-project-pp/product-rpc-service/domain"
	"github.com/private-project-pp/product-rpc-service/entity"
	"gorm.io/gorm"
)

type productWH struct {
	db *gorm.DB
}

func SetupProductWarehouseRepository(db *gorm.DB) domain.ProductWarehouse {
	return &productWH{
		db: db,
	}
}

func (r *productWH) GetProductByID(id string) (out entity.ProductsWarehouse, err error) {

	return out, nil
}

func (r *productWH) AddProduct(in entity.ProductsWarehouse) (err error) {
	return nil
}

func (r *productWH) UpdateProduct(in entity.ProductsWarehouse) (err error) {
	return nil
}
