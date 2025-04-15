package handler

import (
	model "github.com/private-project-pp/pos-grpc-contract/model/product_service"
	"github.com/private-project-pp/product-rpc-service/usecase/product_adm"
)

type productService struct {
	productAdm product_adm.ProductAdministration
	model.UnimplementedProductServiceServer
}

func SetupProductService(
	productAdm product_adm.ProductAdministration,
) model.ProductServiceServer {
	return &productService{
		productAdm: productAdm,
	}
}
