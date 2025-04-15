package entity

import (
	"time"

	"github.com/private-project-pp/product-rpc-service/shared/constant"
)

type UnitOfMeasuremnet struct {
	SecureId  string    `gorm:"column:secure_id"`
	CreatedBy string    `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
	Code      string    `gorm:"column:code"`
	Name      string    `gorm:"column:name"`
}

func (UnitOfMeasuremnet) TableName() string {
	return constant.UnitOfMeasurement
}
