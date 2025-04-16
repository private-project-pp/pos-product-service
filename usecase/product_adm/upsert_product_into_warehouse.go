package product_adm

import (
	"context"

	"github.com/private-project-pp/pos-general-lib/metadata"
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	model "github.com/private-project-pp/pos-grpc-contract/model/product_service"
	"github.com/private-project-pp/product-rpc-service/entity"
	"github.com/private-project-pp/product-rpc-service/shared/utils"
)

func (s *productAdm) UpsertProductIntoWarehouse(ctx context.Context, in *model.UpsertProductWarehouseRequest) (out *model.UpsertProductWarehouseResponse, err error) {
	if in.GetSecureId() == "" {
		id, err := s.AddNewProductIntoWarehouse(ctx, in)
		if err != nil {
			return nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
		}

		out = &model.UpsertProductWarehouseResponse{
			SecureId: id,
		}
		return out, nil
	}

	err = s.UpdateExistingProductInWarehouse(ctx, in)
	if err != nil {
		return nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	out = &model.UpsertProductWarehouseResponse{
		SecureId: in.GetSecureId(),
	}
	return out, nil
}

func (s *productAdm) AddNewProductIntoWarehouse(ctx context.Context, in *model.UpsertProductWarehouseRequest) (out string, err error) {
	product := entity.ProductsWarehouse{
		SecureId:  utils.GenerateUUID(),
		CreatedBy: metadata.GetUAuthUserId(ctx),
		CreatedAt: utils.GetUtcTime(),
		ProductId: in.GetProductId(),
		Amount:    0,
		UnitId:    in.GetUnitId(),
	}

	err = s.productWarehouse.AddProduct(product)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	out = product.SecureId
	return out, nil
}

func (s *productAdm) UpdateExistingProductInWarehouse(ctx context.Context, in *model.UpsertProductWarehouseRequest) (err error) {
	productStock, err := s.productWarehouse.GetProductByID(in.GetSecureId())
	if err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	unitComparator, err := s.productUnitRepo.GetProductUnitByProductAndUnitId(in.GetProductId(), in.GetUnitId(), productStock.UnitId)
	if err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	productStock.UnitId = in.GetUnitId()

	if in.GetRecalculateAmountStock() {

		productStock.Amount = productStock.Amount * unitComparator.Amount
		if unitComparator.SmallerUnitId == in.GetUnitId() {
			productStock.Amount = productStock.Amount / unitComparator.Amount
		}
	}

	err = s.productWarehouse.UpdateProduct(productStock)
	if err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return nil
}
