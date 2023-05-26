package model

import "time"

type Lenders struct {
	Id         int64     `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	Name       string    `gorm:"column:name" validate:"required" json:"name"`
	Limits     string    `gorm:"column:limits" validate:"required" json:"limits"`
	Bunga      string    `gorm:"column:bunga" validate:"required" json:"bunga"`
	Created_At time.Time `gorm:"column:created_at" validate:"required" json:"created_at"`
}
