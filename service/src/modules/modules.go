package modules

import (
	"github.com/fazpass/goliath/v3/module"
	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/ping"
)

type Modules struct {
	Ping         module.ModuleInterface
	Object       module.ModuleInterface
	Organization module.ModuleInterface
}

func Init(app *app.App) *Modules {
	return &Modules{
		Ping:         ping.Init(app),
		Object:       object.Init(app),
		Organization: organization.Init(app),
	}
}
