package response

import (
	"time"

	"github.com/google/uuid"
)

type (
	Organization struct {
		Id           uuid.UUID `json:"id"`
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
