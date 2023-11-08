package response

import (
	"time"
)

type (
	Organization struct {
		Id           uint64    `json:"id"`
		Name         string    `json:"name" required:"true"`
		Email        string    `json:"email" required:"true"`
		Address      string    `json:"address" required:"true"`
		Phone_number string    `json:"phone_number" required:"true"`
		Description  string    `json:"description" required:"true"`
		Logo         string    `json:"logo" required:"true"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)
