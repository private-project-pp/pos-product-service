package entity

import "github.com/private-project-pp/product-rpc-service/shared/constant"

type TransactionDetail struct {
	TransactionId string `gorm:"column:"`
	ProductId     string `gorm:"column:"`
	ProductName   string `gorm:"column:"`
	Amount        string `gorm:"column:"`
	UnitId        string `gorm:"column:"`
	UnitName      string `gorm:"column:"`
	BasePrice     uint64 `gorm:"column:"`
	TotalPrice    uint64 `gorm:"column:"`
}

func (TransactionDetail) TableName() string {
	return constant.TransactionDetail
}
