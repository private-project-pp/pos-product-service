package entity

import (
	"time"

	"github.com/private-project-pp/product-rpc-service/shared/constant"
)

type ProductsImages struct {
	SecureId  string    `gorm:"column:secure_id"`
	CreatedBy string    `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
	ProductId string    `gorm:"column:prodcut_id"`
	FileType  string    `gorm:"column:file_type"`
}

func (ProductsImages) TableName() string {
	return constant.ProductsImages
}
