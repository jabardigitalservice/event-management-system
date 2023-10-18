package repository

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/app"
)

type (
	RepositoryInterface interface {
		GetHealthCheck(context.Context) (map[string]interface{}, error)
	}

	Repository struct {
		app *app.App
	}
)

func Init(app *app.App) *Repository {
	return &Repository{
		app: app,
	}
}
