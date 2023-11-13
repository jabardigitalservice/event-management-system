package response

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
		ID              uuid.UUID      `json:"id"`
		Name            string         `json:"name"`
		Address         string         `json:"address"`
		Description     string         `json:"description"`
		Banner          pq.StringArray `json:"banner"`
		Logo            string         `json:"logo"`
		SocialMedia     []SocialMedia  `json:"social_media"`
		Organizer       string         `json:"organizer"`
		Status          Status         `json:"status"`
		CreatedAt       time.Time      `json:"created_at"`
		UpdatedAt       time.Time      `json:"updated_at"`
		Province        string         `json:"province" validate:"required"`
		City            string         `json:"city" validate:"required"`
		District        string         `json:"district" validate:"required"`
		Village         string         `json:"village" validate:"required"`
		Google_map      string         `json:"google_map" validate:"required"`
		Organizer_email string         `json:"organizer_email" validate:"required"`
		Organizer_phone string         `json:"organizer_phone" validate:"required"`
		ProvinceId      string         `json:"province_id" validate:"required"`
		CityId          string         `json:"city_id" validate:"required"`
		DistrictId      string         `json:"district_id" validate:"required"`
		VillageId       string         `json:"village_id" validate:"required"`
	}
)
