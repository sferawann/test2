package model

import "time"

type Borrower struct {
	Id           int64     `gorm:"primaryKey;column:id" json:"id"`
	Username     string    `gorm:"column:username" json:"username"`
	Password     string    `gorm:"column:password" json:"password"`
	Name         string    `gorm:"column:name" json:"name"`
	Alamat       string    `gorm:"column:alamat" json:"alamat"`
	Phone_Number string    `gorm:"column:phone_number" json:"phone_number"`
	Created_At   time.Time `gorm:"column:created_at" json:"created_at"`
}
