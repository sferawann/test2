package model

import "time"

type Requirements struct {
	Id           int64     `gorm:"primaryKey;column:id" validate:"required" json:"id"`
	Nik          string    `gorm:"column:nik" validate:"required" json:"nik"`
	Name         string    `gorm:"column:name" validate:"required" json:"name"`
	Alamat       string    `gorm:"column:alamat" validate:"required" json:"alamat"`
	Phone_Number string    `gorm:"column:phone_number" validate:"required" json:"phone_number"`
	Created_At   time.Time `gorm:"column:created_at" validate:"required" json:"created_at"`
}
