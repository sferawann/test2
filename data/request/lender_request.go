package request

import "time"

type CreateLenderRequest struct {
	Name       string    `validate:"required" json:"name"`
	Created_At time.Time `validate:"required" json:"created_at"`
}

type UpdateLenderRequest struct {
	Id           int64     `validate:"required"`
	Name         string    `validate:"required" json:"name"`
	Created_At   time.Time `validate:"required" json:"created_at"`
}
