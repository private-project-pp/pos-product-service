package entity

import (
	"time"

	"github.com/private-project-pp/product-rpc-service/shared/constant"
)

type ProductsUnitComparator struct {
	SecureId      string    `gorm:"column:secure_id"`
	CreatedBy     string    `gorm:"column:created_by"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	ModifiedBy    string    `gorm:"column:modified_by"`
	ModifiedAt    time.Time `gorm:"column:modified_at"`
	ProductId     string    `gorm:"column:product_id"`
	UnitId        string    `gorm:"column:unit_id"`
	Amount        uint64    `gorm:"column:amount"`
	SmallerUnitId string    `gorm:"column:smaller_unit_id"`
}

func (ProductsUnitComparator) TableName() string {
	return constant.ProductsUnitComparator
}
