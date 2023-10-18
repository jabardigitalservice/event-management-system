package log

import "context"

type UseCaseLogInterface interface {
	LogSuccess(ctx context.Context, method string, message string)
	LogError(ctx context.Context, method string, e error)
	LogErrorWithAdditionalInfo(context.Context, string, error, map[string]interface{})
}
