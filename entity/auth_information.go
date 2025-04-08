package entity

import (
	"time"

	"github.com/private-project-pp/user-rpc-service/shared/constant"
)

type UsersAuthInformation struct {
	SecureId  string    `gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy string    `gorm:"column:created_by"`
	UpdatedAt time.Time `gorm:"column:created_at"`
	UpdatedBy string    `gorm:"column:created_by"`
	UserId    string    `gorm:"column:user_id"`
	Password  string    `gorm:"column:password"`
	Username  string    `gorm:"column:username"`
	Status    string    `gorm:"column:status"`
}

func (UsersAuthInformation) TableName() string {
	return constant.AuthInformation
}
