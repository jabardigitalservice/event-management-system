package entity

import (
	"time"

	"github.com/google/uuid"
)

type (
	Organization struct {
		Id          uuid.UUID
		Name        string
		Email       string
		Address     string
		PicPhone    string
		Description string
		Logo        string
		CreatedAt   time.Time
		UpdatedAt   time.Time
		Province    string
		City        string
		District    string
		Village     string
		Google_map  string
		PicName     string
		PicPosition string
		Province_id string
		City_id     string
		District_id string
		Village_id  string
	}
)
