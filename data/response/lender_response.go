package response

import "time"

type LenderResponse struct {
	Id           int64     `validate:"required"`
	Name         string    `validate:"required" json:"name"`
	Created_At   time.Time `validate:"required" json:"created_at"`
}
