package entity

import (
	"time"

	"github.com/lib/pq"
)

type (
	Status string

	Object struct {
		ID          uint64         `json:"id"`
		Name        string         `json:"name"`
		Address     string         `json:"address"`
		Description string         `json:"description"`
		Banner      pq.StringArray `json:"banner"`
		Logo        string         `json:"logo"`
		SocialMedia [][]string     `json:"social_media"`
		Organizer   string         `json:"organizer"`
		Status      Status         `json:"status"`
		CreatedAt   time.Time      `json:"created_at"`
		UpdatedAt   time.Time      `json:"updated_at"`
	}
)
