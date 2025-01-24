package middleware

import (
	"chat-platform-api.com/chat-platform-api/src/interface/middleware_interface"
	"chat-platform-api.com/chat-platform-api/src/type/model/common_model"
	"chat-platform-api.com/chat-platform-api/src/variable/api_variable"
)

func ExecMiddlewares(request *common_model.CustomAPIRequest, middlewares ...middleware_interface.MiddlewareImpl) (err error, code int) {
	for _, middleware := range middlewares {
		if err, code := middleware.Do(request); err != nil {
			return err, code
		}
	}
	return nil, api_variable.STATUS_OK
}
