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
		Id          uint64    `db:"id"`
		Name        string    `db:"name"`
		Email       string    `db:"email"`
		Address     string    `db:"address"`
		PhoneNumber string    `db:"phone_number"`
		Logo        string    `db:"logo"`
		CreatedAt   time.Time `db:"created_at"`
		UpdatedAt   time.Time `db:"updated_at"`
	}
)
