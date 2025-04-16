package postgre

import (
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	"github.com/private-project-pp/product-rpc-service/domain"
	"github.com/private-project-pp/product-rpc-service/entity"
	"gorm.io/gorm"
)

type productUnitRepo struct {
	db *gorm.DB
}

func SetupProductUnitRepo(db *gorm.DB) domain.ProductUnit {
	return &productUnitRepo{
		db: db,
	}
}

func (r *productUnitRepo) GetProductUnitById(in string) (out entity.UnitOfMeasuremnet, err error) {
	err = r.db.Model(&entity.UnitOfMeasuremnet{}).Where("secure_id = ?", in).Scan(&out).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return out, stacktrace.Cascade(err, stacktrace.DATA_NOT_FOUND, err.Error())
		}
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return out, nil
}

func (r *productUnitRepo) AddProductUnit(in entity.UnitOfMeasuremnet) (err error) {
	err = r.db.Create(&in).Error
	if err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return nil
}

func (r *productUnitRepo) UpdateProductUnit(in entity.UnitOfMeasuremnet) (err error) {
	err = r.db.Where("secure_id = ?", in.SecureId).Updates(in).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return stacktrace.Cascade(err, stacktrace.DATA_NOT_FOUND, err.Error())
		}
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return nil

}

func (r *productUnitRepo) GetProductUnitByProductAndUnitId(productId, unitId, olderUnitId string) (out entity.ProductsUnitComparator, err error) {
	cond := "product_id = ? AND ((unit_id = ? AND smaller_unit_id = ?) OR (smaller_unit_id = ? AND unit_id ?))"
	err = r.db.Where(cond, productId, unitId, olderUnitId, olderUnitId, unitId).Scan(&out).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return out, stacktrace.Cascade(err, stacktrace.DATA_NOT_FOUND, err.Error())
		}
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return out, nil
}
