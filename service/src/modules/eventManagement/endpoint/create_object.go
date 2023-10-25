package endpoint

import (
	"context"
	"time"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
)

func (e *Endpoint) CreateObject(ctx context.Context, objData entity.Object) (uint64, entity.Object, time.Time, time.Time, error) {
	id, obj, createdAt, updatedAt, err := e.usecase.CreateObject(ctx, objData)
	if err != nil {
		return 0, entity.Object{}, time.Time{}, time.Time{}, err
	}

	object := entity.Object{
		ID:          id,
		Name:        obj.Name,
		Address:     obj.Address,
		Description: obj.Description,
		Banner:      obj.Banner,
		Logo:        obj.Logo,
		SocialMedia: obj.SocialMedia,
		Organizer:   obj.Organizer,
		Status:      obj.Status,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	return id, object, createdAt, updatedAt, nil
}
