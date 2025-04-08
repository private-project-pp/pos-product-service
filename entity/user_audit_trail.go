package entity

import (
	"time"

	"github.com/private-project-pp/user-rpc-service/shared/constant"
)

type UserAuditTrail struct {
	Id           string
	CreatedAt    time.Time
	UserId       string
	ActivityType string
	Detail       []byte
}

func (UserAuditTrail) TableName() string {
	return constant.UserAuditTrail
}
