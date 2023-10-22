package modules

import (
	"github.com/fazpass/goliath/v3/module"
	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/ping"
)

type Modules struct {
	Ping            module.ModuleInterface
	EventManagement module.ModuleInterface
}

func Init(app *app.App) *Modules {
	return &Modules{
		Ping:            ping.Init(app),
		EventManagement: eventManagement.Init(app),
	}
}
