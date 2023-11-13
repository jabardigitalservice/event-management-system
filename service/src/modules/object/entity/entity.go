package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type (
	SocialMedia struct {
		Name string
		Link string
	}

	Status string

	Object struct {
		ID              uuid.UUID
		Name            string
		Address         string
		Description     string
		Banner          pq.StringArray
		Logo            string
		SocialMedia     []SocialMedia
		Organizer       string
		Status          Status
		CreatedAt       time.Time
		UpdatedAt       time.Time
		Province        string
		City            string
		District        string
		Village         string
		Google_map      string
		Organizer_email string
		Organizer_phone string
		ProvinceId      string
		CityId          string
		DistrictId      string
		VillageId       string
		OrganizationID  uuid.UUID
	}
)
