package response

import "time"

type BorrowerResponse struct {
	Id           int64     `validate:"required"`
	Username     string    `validate:"required" json:"username"`
	Password     string    `validate:"required" json:"password"`
	Name         string    `validate:"required" json:"name"`
	Alamat       string    `validate:"required" json:"alamat"`
	Phone_Number string    `validate:"required" json:"phone_number"`
	Created_At   time.Time `validate:"required" json:"created_at"`
}

type LoginResponse struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}
