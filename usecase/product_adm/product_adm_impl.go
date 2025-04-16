package product_adm

import (
	"github.com/private-project-pp/product-rpc-service/domain"
	"gorm.io/gorm"
)

type productAdm struct {
	productRepo      domain.Product
	productUnitRepo  domain.ProductUnit
	dbTx             *gorm.DB
	productWarehouse domain.ProductWarehouse
}

func SetupProductAdministration(
	productRepo domain.Product,
	productUnitRepo domain.ProductUnit,
	dbTx *gorm.DB,
) ProductAdministration {
	return &productAdm{
		productRepo:     productRepo,
		productUnitRepo: productUnitRepo,
		dbTx:            dbTx,
	}
}
