package handler

import (
	"context"

	"github.com/private-project-pp/pos-general-lib/stacktrace"
	model "github.com/private-project-pp/pos-grpc-contract/model/product_service"
)

func (u productService) AddProduct(ctx context.Context, in *model.AddProductRequest) (out *model.AddProductResponse, err error) {
	out, err = u.productAdm.AddingNewProduct(ctx, in)
	if err != nil {
		return nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return out, nil
}

func (u productService) UpsertUnit(ctx context.Context, in *model.UpsertUnitRequest) (out *model.UpsertUnitResponse, err error) {
	out, err = u.productAdm.UpsertUnit(ctx, in)
	if err != nil {
		return nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return out, nil
}
