package entity

import (
	"time"
)

type (
	Organization struct {
		Id           uint64
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
