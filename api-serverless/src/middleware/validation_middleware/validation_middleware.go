package validation_middleware

import (
	"errors"

	"chat-platform-api.com/chat-platform-api/src/interface/middleware_interface"
	"chat-platform-api.com/chat-platform-api/src/interface/repository_interface"
	"chat-platform-api.com/chat-platform-api/src/middleware/common_middleware"
	"chat-platform-api.com/chat-platform-api/src/tool/jwt_tool"
	"chat-platform-api.com/chat-platform-api/src/type/model/common_model"
	"chat-platform-api.com/chat-platform-api/src/type/model/user_model"
	"chat-platform-api.com/chat-platform-api/src/variable/api_variable"
	"chat-platform-api.com/chat-platform-api/src/variable/auth_variable"
)

var (
	middleware *UserValidationMiddleware
)

func GetUserValidationMiddleware() middleware_interface.MiddlewareImpl {
	if middleware == nil {
		middleware = &UserValidationMiddleware{}
	}
	return middleware
}

type UserValidationMiddleware struct {
	common_middleware.Middleware
	userRepository repository_interface.UserRepositoryImpl
}

func (u UserValidationMiddleware) Do(request *common_model.CustomAPIRequest) (err error, code int) {
	accessToken := request.GetHeader("access_token")
	revokeToken := request.GetHeader("revoke_token")
	// access token 검증
	userData, err := jwt_tool.GetData[user_model.UserData](accessToken, auth_variable.JWT_SECRET_KEY)

	if err != nil { // access token 검증되지 않은 경우
		// revoke token 검증
		userData, err = jwt_tool.GetData[user_model.UserData](revokeToken, auth_variable.JWT_SECRET_KEY)
		if err != nil {
			return errors.New("need signIn"), api_variable.STATUS_UNAUTHORIZED
		}
		accessToken = jwt_tool.GenerateToken(userData, auth_variable.JWT_SECRET_KEY, auth_variable.ACCESS_TOKEN_EXPIRATION)
	}

	request.SetParameters(map[string]any{
		"userData":    userData,
		"accessToken": accessToken,
		"revokeToken": revokeToken,
	})

	return nil, api_variable.STATUS_OK
}
