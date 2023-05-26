package model

import (
	"time"
)

type StatusLoans struct {
	Id            int64        `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	Accept_Status AcceptStatus `gorm:"foreignKey;column:id_transaction" validate:"required" json:"transactions"`
	Created_At    time.Time    `gorm:"column:created_at" validate:"required" json:"created_at"`
}
