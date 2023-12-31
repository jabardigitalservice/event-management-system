package app

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/fazpass/goliath/v3/router"
	"github.com/go-chi/chi"
	"github.com/jabardigitalservice/golog/logger"
	gologlogger "github.com/jabardigitalservice/golog/logger"
	"github.com/jabardigitalservice/super-app-services/event/src/constant"
	"github.com/spf13/viper"
	"go.elastic.co/apm/module/apmhttp"
)

type (
	App struct {
		ctx    context.Context
		router *chi.Mux
		logger *logger.Logger
		db     *DB
	}

	DB struct {
		Master *sql.DB
		Slave  *sql.DB
		UserDB *sql.DB
	}
)

func Init() (*App, error) {
	var ctx = context.Background()

	// Initialize the configuration.
	appConfig, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	// Initialize the logger.
	log := logger.Init()

	// Initialize the database connections.
	masterDB := InitPgsqlMaster(ctx, appConfig)
	slaveDB := InitPgsqlSlave(ctx, appConfig)

	app := &App{
		ctx: ctx,
		router: router.InitChi(router.Config{
			Debug: viper.GetBool("APP_DEBUG"),
		}),
		logger: log,
		db: &DB{
			Master: masterDB,
			Slave:  slaveDB,
		},
	}

	return app, nil
}

func (app *App) Context() context.Context {
	return app.ctx
}

func (app *App) GetHttpRouter() *chi.Mux {
	return app.router
}

func (app *App) GetLogger() *logger.Logger {
	return app.logger
}

func (app *App) GetVersion() string {
	return viper.GetString("APP_VERSION")
}

func (app *App) GetDB() *DB {
	return app.db
}

func (app *App) GetStorageBaseUrl() string {
	return viper.GetString("STORAGE_BASE_URL")
}

func (app *App) SetLogger(module string, method string, err error, additionalInfo map[string]interface{}) {
	app.GetLogger().Error(&gologlogger.LoggerData{
		Category:       gologlogger.LoggerApp,
		Service:        constant.ServiceName,
		Module:         module,
		Method:         method,
		Version:        app.GetVersion(),
		AdditionalInfo: additionalInfo,
	}, err)
}

func (app *App) RunHttp() error {
	var hostname = viper.GetString("APP_HOSTNAME")
	if hostname == "" {
		hostname = "0.0.0.0"
	}

	var port = viper.GetString("APP_PORT")
	if port == "" {
		port = "30008"
	}

	var host = hostname + ":" + port

	app.logger.Info(&logger.LoggerData{
		Category: logger.LoggerApp,
		Service:  constant.ServiceName,
		Method:   "startup",
		Version:  app.GetVersion(),
	}, "running")

	return http.ListenAndServe(host, apmhttp.Wrap(app.router))
}
