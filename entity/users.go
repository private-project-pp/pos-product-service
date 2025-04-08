package entity

import (
	"time"

	"github.com/private-project-pp/product-rpc-service/shared/constant"
)

type Products struct {
	Id              uint64     `gorm:"column:id;primaryKey"`
	SecureId        string     `gorm:"column:secure_id"`
	CreatedAt       time.Time  `gorm:"column:created_at"`
	UpdatedAt       *time.Time `gorm:"column:updated_at"`
	Name            string     `gorm:"column:name"`
	Status          uint32     `gorm:"column:status"`
	Note            string     `gorm:"column:note"`
	PurchasingPrice uint64     `gorm:"column:purchasing_price"`
	SellingPrice    uint64     `gorm:"column:selling_price"`
}

func (Products) TableName() string {
	return constant.ProductsEntity
}
