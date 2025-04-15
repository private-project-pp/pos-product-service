package entity

import (
	"time"

	"github.com/private-project-pp/product-rpc-service/shared/constant"
)

type ProductsOnDisplay struct {
	SecureId   string    `gorm:"column:secure_id"`
	CreatedBy  string    `gorm:"column:created_by"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	ModifiedBy string    `gorm:"column:modified_by"`
	ModifiedAt time.Time `gorm:"column:modified_at"`
	ProductId  string    `gorm:"column:product_id"`
	Price      uint64    `gorm:"column:price"`
	Amount     uint64    `gorm:"column:amount"`
	UnitId     string    `gorm:"column:unit_id"`
}

func (ProductsOnDisplay) TableName() string {
	return constant.ProductsOnDisplay
}
