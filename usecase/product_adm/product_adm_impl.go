package product_adm

import "github.com/private-project-pp/product-rpc-service/domain"

type productAdm struct {
	productRepo     domain.Product
	productUnitRepo domain.ProductUnit
}

func SetupProductAdministration(
	productRepo domain.Product,
	productUnitRepo domain.ProductUnit,
) ProductAdministration {
	return &productAdm{
		productRepo:     productRepo,
		productUnitRepo: productUnitRepo,
	}
}
