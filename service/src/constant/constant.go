package constant

type contextKey string

const (
	HeaderSessionIDKey = "X-Session-ID"
	HeaderClientIDKey  = "X-Client-ID"
	HeaderUserIDKey    = "X-User-ID"
	HeaderUserName     = "X-User-Name"

	ServiceName = "event"

	ModuleNameEventManagement = "event-management"

	ContextKeyToken    contextKey = "ck:token"
	ContextKeyUserId   contextKey = "ck:user-id"
	ContextKeyUserName contextKey = "ck:user-name"
	ContextFCMToken    contextKey = "ck:fcm-token"

	LogCategoryApp      = "app"
	LogCategoryUsecase  = "usecase"
	LogCategoryExternal = "external"
	LogCategoryRouter   = "router"
)
