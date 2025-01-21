package api_middleware

import (
	"errors"
	"strings"

	"chat-platform-api.com/chat-platform-api/src/type/model/common_model"
	api_variable "chat-platform-api.com/chat-platform-api/src/variable/api_varialbe"
)

var (
	postURLS   = []string{"platform/register", "platform/issueToken", "platform/update", "platform/withdraw"}
	getURLS    = []string{"platform/getData"}
	deleteURLS = []string{"platform/delete"}
)

func CheckAPIUrlMiddleware(request *common_model.CustomAPIRequest) error {
	isContains := func(allowUrls []string, url string) bool {
		for _, allow := range allowUrls {
			if strings.Contains(allow, url) {
				return true
			}
		}
		return false
	}

	urls := strings.Split(request.Path, "/")

	if len(urls) < 3 || urls[1] != "api" {
		return errors.New(api_variable.RESPONSE_INVALID_PATH)
	}

	mainUrl := strings.Join(urls[2:], "/")
	allowMethod := ""

	if isContains(postURLS, mainUrl) {
		allowMethod = "POST"
	} else if isContains(getURLS, mainUrl) {
		allowMethod = "GET"
	} else if isContains(deleteURLS, mainUrl) {
		allowMethod = "DELETE"
	}

	request.HTTPMethod = strings.ToUpper(request.HTTPMethod)

	if allowMethod == "" || request.HTTPMethod != allowMethod {
		return errors.New(api_variable.RESPONSE_INVALID_PATH)
	}

	request.GlobalParameter["url"] = mainUrl
	

	return nil
}
