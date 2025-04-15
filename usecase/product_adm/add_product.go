package product_adm

import (
	"context"

	"github.com/private-project-pp/pos-general-lib/metadata"
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	model "github.com/private-project-pp/pos-grpc-contract/model/product_service"
	"github.com/private-project-pp/product-rpc-service/entity"
	"github.com/private-project-pp/product-rpc-service/shared/utils"
)

func (s productAdm) AddingNewProduct(ctx context.Context, in *model.AddProductRequest) (out *model.AddProductResponse, err error) {
	var productImages []entity.ProductsImages
	product := entity.Product{
		SecureId:  utils.GenerateUUID(),
		CreatedBy: metadata.GetUAuthUserId(ctx),
		CreatedAt: utils.GetUtcTime(),
		Name:      in.GetName(),
		Barcode:   in.GetBarcode(),
		Status:    in.GetStatus(),
		Note:      in.GetNote(),
		ImageId:   in.GetImage(),
		UnitId:    in.GetUnit(),
	}

	for _, image := range in.GetImages() {
		productImages = append(productImages, entity.ProductsImages{
			SecureId:  image.GetSecureId(),
			CreatedBy: metadata.GetUAuthUserId(ctx),
			CreatedAt: utils.GetUtcTime(),
			FileName:  image.GetFileName(),
			FileType:  image.GetFileType(),
		})
	}

	err = s.productRepo.AddProduct(product)
	if err != nil {
		return nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	err = s.productRepo.BulkAddProductImages(productImages)
	if err != nil {
		return nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	out = &model.AddProductResponse{
		SecureId: product.SecureId,
	}

	return out, nil
}
