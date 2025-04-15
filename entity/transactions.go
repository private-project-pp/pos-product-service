package entity

import "github.com/private-project-pp/product-rpc-service/shared/constant"

type Transactions struct {
	SecureId               string `gorm:"column:secure_id"`
	CreatedBy              string `gorm:"column:created_by"`
	CreatedAt              string `gorm:"column:created_at"`
	TotalTransaction       uint64 `gorm:"column:total_transaction"`
	BuyerName              string `gorm:"column:buyer_name"`
	PaymentStatus          uint8  `gorm:"column:status"`
	PaymentMethodId        string `gorm:"column:payment_method_id"`
	GeneratedinvoiceNumber string `gorm:"column:generated_invoice_number"`
}

func (Transactions) TableName() string {
	return constant.Transactions
}
