package main

import (
	"log"
	"net/http"

	helperresponse "github.com/fazpass/goliath/v3/helper/http/response"
	middlewaregoliath "github.com/fazpass/goliath/v3/middleware"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	gologconstant "github.com/jabardigitalservice/golog/constant"
	"github.com/jabardigitalservice/super-app-services/event/src/app"
	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	"github.com/jabardigitalservice/super-app-services/event/src/modules"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"

	middlewareheadertoctx "github.com/jabardigitalservice/utilities-go/middleware/http/headers-to-ctx"
	"github.com/spf13/viper"
)

func main() {
	app, initErr := app.Init()

	if initErr != nil {
		log.Panic(initErr)
	}

	router := app.GetHttpRouter()
	module := modules.Init(app)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		// ExposedHeaders:   []string{"Link"},
		// AllowCredentials: false,
		// MaxAge:           300,
	}))
	router.Use(middleware.CleanPath)
	router.Use(middleware.RealIP)
	router.Use(middleware.Heartbeat("/"))
	router.Use(middlewareheadertoctx.Mapping(map[string]interface{}{
		gologconstant.CtxSessionIDKey: constant.HeaderSessionIDKey,
		gologconstant.CtxClientIDKey:  constant.HeaderClientIDKey,
		gologconstant.CtxUserIDKey:    constant.HeaderUserIDKey,
	}))

	router.Use(middlewaregoliath.Recoverer(middlewaregoliath.RecovererOptions{
		Debug: viper.GetBool("APP_DEBUG"),
		Response: middlewaregoliath.RecovererResponse{
			Status: http.StatusInternalServerError,
			Body: helperresponse.Response{
				Status:  false,
				Message: "Error",
				Code:    httpresponse.InternalServerError,
			},
		},
	}))

	router.Mount("/v1/event/ping", module.Ping.GetHttpRouter())
	router.Mount("/v1/event/object", module.Object.GetHttpRouter())
	router.Mount("/v1/event/organization", module.Organization.GetHttpRouter())

	var err = app.RunHttp()
	if err != nil {
		log.Panic(err)
	}
}
