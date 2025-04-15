package entity

import (
	"time"

	"github.com/private-project-pp/product-rpc-service/shared/constant"
)

type Product struct {
	Id          uint64     `gorm:"column:id;primaryKey"`
	SecureId    string     `gorm:"column:secure_id"`
	CreatedBy   string     `gorm:"column:created_by"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	ModifieddBy string     `gorm:"column:modified_by"`
	ModifiedAt  *time.Time `gorm:"column:modified_at"`
	Name        string     `gorm:"column:name"`
	Barcode     string     `gorm:"column:barcode"`
	Status      uint32     `gorm:"column:status"`
	Note        string     `gorm:"column:note"`
	ImageId     string     `gorm:"column:image_id"`
	UnitId      string     `gorm:"column:unit_id"`
}

func (Product) TableName() string {
	return constant.ProductsEntity
}
