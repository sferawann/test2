package model

import "time"

type Lender struct {
	Id         int64     `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	Name       string    `gorm:"column:name" validate:"required" json:"name"`
	Created_At time.Time `gorm:"column:created_at" validate:"required" json:"created_at"`
}
