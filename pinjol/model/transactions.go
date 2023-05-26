package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transactions struct {
	Id                int64           `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	Requirements      Requirements    `gorm:"foreignKey;column:id_requirement" validate:"required" json:"requirements"`
	Lenders           Lenders         `gorm:"foreignKey;column:id_lender" validate:"required" json:"Lenders"`
	Users             Users           `gorm:"foreignKey;column:id_user" validate:"required" json:"users"`
	Amount            decimal.Decimal `gorm:"column:amount" validate:"required" json:"amount"`
	Transactions_Date time.Time       `gorm:"column:transaction_date" validate:"required" json:"transaction_date"`
	Due_Date          time.Time       `gorm:"column:due_date" validate:"required" json:"due_date"`
}
