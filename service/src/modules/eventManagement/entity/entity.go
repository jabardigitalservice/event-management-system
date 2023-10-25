package entity

import (
	"time"

	"github.com/lib/pq"
)

type (
	Organization struct {
		Id          uint64    `db:"id"`
		Name        string    `db:"name"`
		Email       string    `db:"email"`
		Address     string    `db:"address"`
		PhoneNumber string    `db:"phone_number"`
		Logo        string    `db:"logo"`
		CreatedAt   time.Time `db:"created_at"`
		UpdatedAt   time.Time `db:"updated_at"`
	}

	SocialMedia struct {
		Name string `json:"name"`
		Link string `json:"link"`
	}

	Status string

	Object struct {
		ID          uint64         `json:"id"`
		Name        string         `json:"name"`
		Address     string         `json:"address"`
		Description string         `json:"description"`
		Banner      pq.StringArray `json:"banner"`
		Logo        string         `json:"logo"`
		SocialMedia []SocialMedia  `json:"social_media"`
		Organizer   string         `json:"organizer"`
		Status      Status         `json:"status"`
		CreatedAt   time.Time      `json:"created_at"`
		UpdatedAt   time.Time      `json:"updated_at"`
	}
)
