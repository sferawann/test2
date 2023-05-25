package model

import "time"

type LoanProduct struct {
	Id          int64     `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	Name        string    `gorm:"column:name" validate:"required" json:"name"`
	Description string    `gorm:"column:description" validate:"required" json:"description"`
	Persyaratan string    `gorm:"column:persyaratan" validate:"required" json:"persyaratan"`
	Created_At  time.Time `gorm:"column:created_at" validate:"required" json:"created_at"`
}
