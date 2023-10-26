package endpoint

import (
	"context"
	"time"

	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/usecase"
)

type (
	EndpointInterface interface {
		CreateOrganizations(ctx context.Context, orgData entity.Organization) (uint64, entity.Organization, error)
		CreateObject(ctx context.Context, objData entity.Object) (uint64, entity.Object, time.Time, time.Time, error)
		GetObjects(ctx context.Context, page int, perPage int) ([]entity.Object, error)
		GetObjectByID(ctx context.Context, objectID uint64) (entity.Object, error)
		UpdateObject(ctx context.Context, obj entity.Object) (entity.Object, error)
		DeleteObject(ctx context.Context, id uint64) error
	}

	Endpoint struct {
		app     *app.App
		usecase usecase.UsecaseInterface
	}
)

func Init(app *app.App, usecase usecase.UsecaseInterface) EndpointInterface {
	return &Endpoint{
		app:     app,
		usecase: usecase,
	}
}

type (
	Organization struct {
		Id          uint64    `json:"id"`
		Name        string    `json:"name"`
		Email       string    `json:"email"`
		Address     string    `json:"address"`
		PhoneNumber string    `json:"phone_number"`
		Logo        string    `json:"logo"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
