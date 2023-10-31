package app

import (
	"context"

	gologconstant "github.com/jabardigitalservice/golog/constant"
	"github.com/jabardigitalservice/golog/logger"
	"github.com/jabardigitalservice/super-app-services/event/src/constant"
)

func (app *App) LogSuccess(ctx context.Context, method string, message string) {
	data := &logger.LoggerData{
		Category: logger.LoggerApp,
		Service:  constant.ServiceName,
		Module:   constant.ModuleNameobject,
		Method:   method,
		Version:  app.GetVersion(),
	}

	if ctx.Value(gologconstant.CtxUserIDKey) != nil {
		data.UserID = ctx.Value(gologconstant.CtxUserIDKey).(string)
	}

	if ctx.Value(gologconstant.CtxSessionIDKey) != nil {
		data.SessionID = ctx.Value(gologconstant.CtxSessionIDKey).(string)
	}

	if ctx.Value(gologconstant.CtxClientIDKey) != nil {
		data.ClientID = ctx.Value(gologconstant.CtxClientIDKey).(string)
	}

	app.logger.Info(data, message)

}

func (app *App) LogError(ctx context.Context, method string, e error) {
	data := &logger.LoggerData{
		Category: logger.LoggerApp,
		Service:  constant.ServiceName,
		Module:   constant.ModuleNameobject,
		Method:   method,
		Version:  app.GetVersion(),
	}

	if ctx.Value(gologconstant.CtxUserIDKey) != nil {
		data.UserID = ctx.Value(gologconstant.CtxUserIDKey).(string)
	}

	if ctx.Value(gologconstant.CtxSessionIDKey) != nil {
		data.SessionID = ctx.Value(gologconstant.CtxSessionIDKey).(string)
	}

	if ctx.Value(gologconstant.CtxClientIDKey) != nil {
		data.ClientID = ctx.Value(gologconstant.CtxClientIDKey).(string)
	}

	app.logger.Error(data, e)
}

func (app *App) LogErrorWithAdditionalInfo(ctx context.Context, method string, e error, additionalInfo map[string]interface{}) {
	data := &logger.LoggerData{
		Category:       logger.LoggerApp,
		Service:        constant.ServiceName,
		Module:         constant.ModuleNameobject,
		Method:         method,
		Version:        app.GetVersion(),
		AdditionalInfo: additionalInfo,
	}

	if ctx.Value(gologconstant.CtxUserIDKey) != nil {
		data.UserID = ctx.Value(gologconstant.CtxUserIDKey).(string)
	}

	if ctx.Value(gologconstant.CtxSessionIDKey) != nil {
		data.SessionID = ctx.Value(gologconstant.CtxSessionIDKey).(string)
	}

	if ctx.Value(gologconstant.CtxClientIDKey) != nil {
		data.ClientID = ctx.Value(gologconstant.CtxClientIDKey).(string)
	}

	app.logger.Error(data, e)
}

func (app *App) LogSuccessWithAdditionalInfo(ctx context.Context, method string, message string, additionalInfo map[string]interface{}) {
	data := &logger.LoggerData{
		Category:       logger.LoggerApp,
		Service:        constant.ServiceName,
		Module:         constant.ModuleNameobject,
		Method:         method,
		Version:        app.GetVersion(),
		AdditionalInfo: additionalInfo,
	}

	if ctx.Value(gologconstant.CtxUserIDKey) != nil {
		data.UserID = ctx.Value(gologconstant.CtxUserIDKey).(string)
	}

	if ctx.Value(gologconstant.CtxSessionIDKey) != nil {
		data.SessionID = ctx.Value(gologconstant.CtxSessionIDKey).(string)
	}

	if ctx.Value(gologconstant.CtxClientIDKey) != nil {
		data.ClientID = ctx.Value(gologconstant.CtxClientIDKey).(string)
	}

	app.logger.Info(data, message)
}

type LogData struct {
	Service        string
	Category       string
	Module         string
	Method         string
	Message        string
	Error          error
	AdditionalInfo map[string]interface{}
}

func (app *App) SetLoggerV2(ctx context.Context, logData *LogData) {
	data := &logger.LoggerData{
		Service: logData.Service,
		Module:  logData.Module,
		Method:  logData.Method,
		Version: app.GetVersion(),
	}

	switch logData.Category {
	case constant.LogCategoryApp:
		data.Category = logger.LoggerApp
	case constant.LogCategoryRouter:
		data.Category = logger.LoggerRouter
	case constant.LogCategoryUsecase:
		data.Category = logger.LoggerUsecase
	case constant.LogCategoryExternal:
		data.Category = logger.LoggerExternal
	default:
		data.Category = logger.LoggerApp
	}

	if ctx.Value(gologconstant.CtxUserIDKey) != nil {
		data.UserID = ctx.Value(gologconstant.CtxUserIDKey).(string)
	}

	if ctx.Value(gologconstant.CtxSessionIDKey) != nil {
		data.SessionID = ctx.Value(gologconstant.CtxSessionIDKey).(string)
	}

	if ctx.Value(gologconstant.CtxClientIDKey) != nil {
		data.ClientID = ctx.Value(gologconstant.CtxClientIDKey).(string)
	}

	if logData.Error != nil {
		app.logger.Error(data, logData.Error)
	} else {
		app.logger.Info(data, logData.Message)
	}
}
