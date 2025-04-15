package product_adm

import (
	"context"

	"github.com/private-project-pp/pos-general-lib/metadata"
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	model "github.com/private-project-pp/pos-grpc-contract/model/product_service"
	"github.com/private-project-pp/product-rpc-service/entity"
	"github.com/private-project-pp/product-rpc-service/shared/utils"
)

func (s *productAdm) UpsertUnit(ctx context.Context, in *model.UpsertUnitRequest) (out *model.UpsertUnitResponse, err error) {
	if in.GetSecureId() == "" {
		secureId, err := s.AddNewProductUnit(ctx, in)
		if err != nil {
			return nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
		}

		out = &model.UpsertUnitResponse{
			SecureId: secureId,
		}
		return out, nil
	}

	err = s.UpdateExistingUnit(in)
	if err != nil {
		return nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	out = &model.UpsertUnitResponse{
		SecureId: in.GetSecureId(),
	}

	return out, nil
}

func (s *productAdm) AddNewProductUnit(ctx context.Context, in *model.UpsertUnitRequest) (out string, err error) {
	newUnit := entity.UnitOfMeasuremnet{
		SecureId:  utils.GenerateUUID(),
		CreatedBy: metadata.GetUAuthUserId(ctx),
		CreatedAt: utils.GetUtcTime(),
		Code:      in.GetCode(),
		Name:      in.GetName(),
	}

	err = s.productUnitRepo.AddProductUnit(newUnit)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	out = newUnit.SecureId
	return out, nil
}

func (s *productAdm) UpdateExistingUnit(in *model.UpsertUnitRequest) (err error) {
	unit, err := s.productUnitRepo.GetProductUnitById(in.GetSecureId())
	if err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	unit.Name = in.GetName()
	unit.Code = in.GetCode()

	err = s.productUnitRepo.UpdateProductUnit(unit)
	if err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return nil
}
