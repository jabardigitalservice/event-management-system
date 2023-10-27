package request

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type (
	SocialMedia struct {
		Name string `json:"name"`
		Link string `json:"link"`
	}

	Status string

	Object struct {
		ID          uuid.UUID      `json:"id" validate:"required"`
		Name        string         `json:"name" validate:"required"`
		Address     string         `json:"address" validate:"required"`
		Description string         `json:"description" validate:"required"`
		Banner      pq.StringArray `json:"banner"`
		Logo        string         `json:"logo"`
		SocialMedia []SocialMedia  `json:"social_media"`
		Organizer   string         `json:"organizer" validate:"required"`
		Status      Status         `json:"status" validate:"required"`
		CreatedAt   time.Time      `json:"created_at"`
		UpdatedAt   time.Time      `json:"updated_at"`
	}
)
