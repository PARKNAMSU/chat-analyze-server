package api_middleware

import (
	"errors"
	"strings"

	"chat-platform-api.com/chat-platform-api/src/interface/middleware_interface"
	"chat-platform-api.com/chat-platform-api/src/middleware/common_middleware"
	"chat-platform-api.com/chat-platform-api/src/type/model/common_model"
	"chat-platform-api.com/chat-platform-api/src/variable/api_variable"
)

var (
	middleware *APIMiddleware
)

func GetAPIMiddleware() middleware_interface.MiddlewareImpl {
	if middleware == nil {
		middleware = &APIMiddleware{}
	}
	return middleware
}

type APIMiddleware struct {
	common_middleware.Middleware
}

func (*APIMiddleware) Do(request *common_model.CustomAPIRequest) (error, int) {
	urls := strings.Split(request.Path, "/")

	if len(urls) < 3 || urls[1] != "api" {
		return errors.New(api_variable.RESPONSE_INVALID_PATH), api_variable.STATUS_BAD_REQUEST
	}

	mainUrl := strings.Join(urls[2:], "/")
	request.SetParameter("url", mainUrl)

	return nil, api_variable.STATUS_OK
}
