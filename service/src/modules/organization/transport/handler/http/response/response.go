package response

import (
	"time"

	"github.com/google/uuid"
)

type (
	Organization struct {
		Id          uuid.UUID `json:"id"`
		Name        string    `json:"name" required:"true"`
		Email       string    `json:"email" required:"true"`
		Address     string    `json:"address" required:"true"`
		Description string    `json:"description" required:"true"`
		Logo        string    `json:"logo" required:"true"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Province    string    `json:"province" validate:"required"`
		City        string    `json:"city" validate:"required"`
		District    string    `json:"district" validate:"required"`
		Village     string    `json:"village" validate:"required"`
		Google_map  string    `json:"google_map" validate:"required"`
		PicName     string    `json:"pic_name" validate:"required"`
		PicPosition string    `json:"pic_position" validate:"required"`
		PicPhone    string    `json:"pic_phone" required:"true"`
		ProvinceId  string    `json:"province_id" validate:"required"`
		CityId      string    `json:"city_id" validate:"required"`
		DistrictId  string    `json:"district_id" validate:"required"`
		VillageId   string    `json:"village_id" validate:"required"`
	}
)