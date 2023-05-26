package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Payments struct {
	Id             int64           `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	Transactions   Transactions    `gorm:"foreignKey;column:id_transaction" validate:"required" json:"transactions"`
	Payment_Method PaymentMethods  `gorm:"foreignKey;column:id_payment_method" validate:"required" json:"payment_method"`
	Payment_Amount decimal.Decimal `gorm:"column:payment_amount" validate:"required" json:"payment_amount"`
	Payment_Date   time.Time       `gorm:"column:payment_date" validate:"required" json:"payment_date"`
}
