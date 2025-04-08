package product_adm

import "github.com/private-project-pp/product-rpc-service/domain"

type productAdm struct {
	productRepo domain.Product
}

func SetupProductAdministration(
	productRepo domain.Product,
) ProductAdministration {
	return &productAdm{
		productRepo: productRepo,
	}
}
