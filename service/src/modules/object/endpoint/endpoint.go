package endpoint

import (
	"context"
	"time"

	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/usecase"
)

type (
	EndpointInterface interface {
		CreateObject(ctx context.Context, objData entity.Object) (uint64, entity.Object, time.Time, time.Time, error)
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
