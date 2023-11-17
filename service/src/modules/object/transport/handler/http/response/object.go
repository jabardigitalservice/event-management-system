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
		ID               uuid.UUID      `json:"id"`
		Name             string         `json:"name"`
		Address          string         `json:"address"`
		Description      string         `json:"description"`
		Banner           pq.StringArray `json:"banner"`
		Logo             string         `json:"logo"`
		SocialMedia      []SocialMedia  `json:"social_media"`
		Organizer        string         `json:"organizer"`
		Status           Status         `json:"status"`
		CreatedAt        time.Time      `json:"created_at"`
		UpdatedAt        time.Time      `json:"updated_at"`
		Province         string         `json:"province" `
		City             string         `json:"city" `
		District         string         `json:"district" `
		Village          string         `json:"village" `
		GoogleMap        string         `json:"google_map" `
		OrganizerEmail   string         `json:"organizer_email" `
		OrganizerPhone   string         `json:"organizer_phone" `
		ProvinceId       string         `json:"province_id" `
		CityID           string         `json:"city_id" `
		DistrictID       string         `json:"district_id" `
		VillageID        string         `json:"village_id" `
		OrganizationID   uuid.UUID      `json:"organization_id" `
		OrganizationName string         `json:"organization_name" `
	}
)
