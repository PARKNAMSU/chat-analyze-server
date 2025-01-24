package middleware_interface

import "chat-platform-api.com/chat-platform-api/src/type/model/common_model"

type MiddlewareImpl interface {
	Do(request *common_model.CustomAPIRequest) (err error, code int)
}
