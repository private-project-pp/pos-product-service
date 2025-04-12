package entity

import (
	"time"

	"github.com/private-project-pp/product-rpc-service/shared/constant"
)

type ProductsWarehouse struct {
	SecureId   string    `gorm:"column:secure_id"`
	CreatedBy  string    `gorm:"column:created_by"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	ModifiedBy string    `gorm:"column:modified_by"`
	ModifiedAt string    `gorm:"column:modified_by"`
	ProductId  string    `gorm:"column:product_id"`
	Amount     uint64    `gorm:"column:amount"`
	UnitId     string    `gorm:"column:unit_id"`
}

func (ProductsWarehouse) TableName() string {
	return constant.ProductsWarehouse
}
