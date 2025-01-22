package api_middleware

import (
	"errors"
	"strings"

	"chat-platform-api.com/chat-platform-api/src/type/model/common_model"
	"chat-platform-api.com/chat-platform-api/src/variable/api_variable"
)

func CheckAPIUrlMiddleware(request *common_model.CustomAPIRequest) error {
	urls := strings.Split(request.Path, "/")

	if len(urls) < 3 || urls[1] != "api" {
		return errors.New(api_variable.RESPONSE_INVALID_PATH)
	}

	mainUrl := strings.Join(urls[2:], "/")
	request.GlobalParameter["url"] = mainUrl

	return nil
}
