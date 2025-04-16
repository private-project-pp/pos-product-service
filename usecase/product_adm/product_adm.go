package product_adm

import (
	"context"

	model "github.com/private-project-pp/pos-grpc-contract/model/product_service"
)

type ProductAdministration interface {
	UpsertUnit(ctx context.Context, in *model.UpsertUnitRequest) (out *model.UpsertUnitResponse, err error)
	AddNewProductUnit(ctx context.Context, in *model.UpsertUnitRequest) (out string, err error)
	UpdateExistingUnit(in *model.UpsertUnitRequest) (err error)
	AddingNewProduct(ctx context.Context, in *model.AddProductRequest) (out *model.AddProductResponse, err error)
	UpsertProductIntoWarehouse(ctx context.Context, in *model.UpsertProductWarehouseRequest) (out *model.UpsertProductWarehouseResponse, err error)
	AddNewProductIntoWarehouse(ctx context.Context, in *model.UpsertProductWarehouseRequest) (out string, err error)
	UpdateExistingProductInWarehouse(ctx context.Context, in *model.UpsertProductWarehouseRequest) (err error)
}
