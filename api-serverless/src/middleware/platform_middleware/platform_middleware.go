package platform_middleware

import (
	"chat-platform-api.com/chat-platform-api/src/interface/middleware_interface"
	"chat-platform-api.com/chat-platform-api/src/interface/repository_interface/platform_interface"
	"chat-platform-api.com/chat-platform-api/src/middleware/common_middleware"
	"chat-platform-api.com/chat-platform-api/src/type/model/common_model"
)

var (
	middleware *PlatformMiddleware
)

func GetUserValidationMiddleware() middleware_interface.MiddlewareImpl {
	if middleware == nil {
		middleware = &PlatformMiddleware{}
	}
	return middleware
}

type PlatformMiddleware struct {
	common_middleware.Middleware
	platformRepo platform_interface.PlatformRepositoryImpl
}

// Do implements middleware_interface.MiddlewareImpl.
func (p *PlatformMiddleware) Do(request *common_model.CustomAPIRequest) (err error, code int) {
	panic("unimplemented")
}
