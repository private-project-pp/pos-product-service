package postgre

import (
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	"github.com/private-project-pp/product-rpc-service/domain"
	"github.com/private-project-pp/product-rpc-service/entity"
	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func SetupProductsRepo(db *gorm.DB) domain.Product {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) AddProduct(tx *gorm.DB, in entity.Product) (err error) {
	err = r.db.Create(&in).Error
	if err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return nil
}

func (r *productRepo) BulkAddProductImages(tx *gorm.DB, in []entity.ProductsImages) (err error) {
	err = r.db.Create(&in).Error
	if err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return nil
}
