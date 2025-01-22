package middleware

import "chat-platform-api.com/chat-platform-api/src/type/model/common_model"

func ExecMiddlewares(request *common_model.CustomAPIRequest, middlewares ...common_model.Middleware) error {
	for _, middleware := range middlewares {
		if err := middleware(request); err != nil {
			return err
		}
	}
	return nil
}