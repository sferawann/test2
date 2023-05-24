package model

import "time"

type Lender struct {
	Id         int64     `gorm:"primaryKey;column:id" JSON:"id"`
	Name       string    `gorm:"column:name" JSON:"name"`
	Created_At time.Time `gorm:"column:created_at" JSON:"created_at"`
}
