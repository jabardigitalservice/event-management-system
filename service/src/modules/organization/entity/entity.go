package entity

import (
	"time"

	"github.com/google/uuid"
)

type (
	Organization struct {
		Id           uuid.UUID
		Name         string
		Email        string
		Address      string
		Phone_number string
		Description  string
		Logo         string
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
)
