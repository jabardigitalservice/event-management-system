package response

import (
	"time"

	"github.com/google/uuid"
)

type (
	Organization struct {
		Id          uuid.UUID `json:"id"`
		Name        string    `json:"name" `
		Email       string    `json:"email" `
		Address     string    `json:"address" `
		Description string    `json:"description" `
		Logo        string    `json:"logo" `
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Province    string    `json:"province" `
		City        string    `json:"city" `
		District    string    `json:"district" `
		Village     string    `json:"village" `
		GoogleMap   string    `json:"google_map" `
		PicName     string    `json:"pic_name" `
		PicPosition string    `json:"pic_position" `
		PicPhone    string    `json:"pic_phone" `
		ProvinceID  string    `json:"province_id"`
		CityID      string    `json:"city_id" `
		DistrictID  string    `json:"district_id" `
		VillageID   string    `json:"village_id" `
	}
)
