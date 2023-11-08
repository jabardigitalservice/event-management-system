package app

import (
	"context"

	gologconstant "github.com/jabardigitalservice/golog/constant"
	"github.com/jabardigitalservice/golog/logger"
	"github.com/jabardigitalservice/super-app-services/event/src/constant"
)

type AppLogInterface interface {
	Log(ctx context.Context, module string, category string) *AppLogger
	Success(ctx context.Context, category string, method string, message string)
	Error(ctx context.Context, category string, method string, err error)
}

type AppLogger struct {
	Logger *logger.Logger
	Data   logger.LoggerData
}

func InitAppLogger(app *Config, service string) *AppLogger {
	appLogger := &AppLogger{
		Logger: logger.Init(),
		Data: logger.LoggerData{
			Service: service,
			Version: app.AppVersion,
		},
	}

	return appLogger
}

func (appLogger *AppLogger) Log(ctx context.Context, category string, module string) *AppLogger {
	appLogger.Data.Module = module

	switch category {
	case constant.LogCategoryApp:
		appLogger.Data.Category = logger.LoggerApp
	case constant.LogCategoryUsecase:
		appLogger.Data.Category = logger.LoggerUsecase
	case constant.LogCategoryRouter:
		appLogger.Data.Category = logger.LoggerRouter
	case constant.LogCategoryExternal:
		appLogger.Data.Category = logger.LoggerExternal
	default:
		appLogger.Data.Category = logger.LoggerApp
	}

	if ctx.Value(gologconstant.CtxUserIDKey) != nil {
		appLogger.Data.UserID = ctx.Value(gologconstant.CtxUserIDKey).(string)
	}

	if ctx.Value(gologconstant.CtxSessionIDKey) != nil {
		appLogger.Data.SessionID = ctx.Value(gologconstant.CtxSessionIDKey).(string)
	}

	if ctx.Value(gologconstant.CtxClientIDKey) != nil {
		appLogger.Data.ClientID = ctx.Value(gologconstant.CtxClientIDKey).(string)
	}

	return appLogger
}

func (appLogger *AppLogger) WithAdditionalInfo(info map[string]interface{}) *AppLogger {
	appLogger.Data.AdditionalInfo = info

	return appLogger
}

func (appLogger *AppLogger) Success(method string, message string) {
	appLogger.Data.Method = method

	appLogger.Logger.Info(&appLogger.Data, message)
}

func (appLogger *AppLogger) Error(method string, err error) {
	appLogger.Data.Method = method

	appLogger.Logger.Error(&appLogger.Data, err)
}
