package entity

import (
	"time"

	"github.com/lib/pq"
)

type (
	SocialMedia struct {
		Name string
		Link string
	}

	Status string

	Object struct {
		ID          uint64
		Name        string
		Address     string
		Description string
		Banner      pq.StringArray
		Logo        string
		SocialMedia []SocialMedia
		Organizer   string
		Status      Status
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
