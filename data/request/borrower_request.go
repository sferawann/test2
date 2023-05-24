package request

import "time"

type CreateBorrowerRequest struct {
	Username     string    `validate:"required" json:"username"`
	Password     string    `validate:"required" json:"password"`
	Name         string    `validate:"required" json:"name"`
	Alamat       string    `validate:"required" json:"alamat"`
	Phone_Number string    `validate:"required" json:"phone_number"`
	Created_At   time.Time `validate:"required" json:"created_at"`
}

type UpdateBorrowerRequest struct {
	Id           int64     `validate:"required"`
	Username     string    `validate:"required" json:"username"`
	Password     string    `validate:"required" json:"password"`
	Name         string    `validate:"required" json:"name"`
	Alamat       string    `validate:"required" json:"alamat"`
	Phone_Number string    `validate:"required" json:"phone_number"`
	Created_At   time.Time `validate:"required" json:"created_at"`
}

type LoginRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}
