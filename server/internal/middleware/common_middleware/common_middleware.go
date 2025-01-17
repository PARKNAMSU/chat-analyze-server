package common_middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"chat-analyze.com/chat-analyze-server/internal/option"
	"chat-analyze.com/chat-analyze-server/internal/tools"
	"chat-analyze.com/chat-analyze-server/internal/usecase/platform_usecase"
)

var (
	platformUsecase = platform_usecase.GetUseCase()
)

// 공통으로 헤더에 값 설정하는 미들웨어
func SetHeader(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Content-Type", "application/json")
	next.ServeHTTP(w, r)
}

// api key 검증 미들웨어
func APIKeyValidation(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	apiKey := r.Header.Get("x-api-key")
	if apiKey != os.Getenv("SERVER_API_KEY") {
		tools.SendErrorResponse(w, option.INVALID_API_KEY, http.StatusUnauthorized)
		return
	}
	next(w, r)
}

func PlatformValidation(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	tokens := strings.Split(r.Header.Get("authorization"), " ")
	hostname := r.Header.Get("hostname")

	if len(tokens) != 2 || tokens[0] != "Bearer" {
		tools.SendErrorResponse(w, option.INVALID_AUTHORIZATION, http.StatusUnauthorized)
		return
	}

	platform, err := platformUsecase.APIKeyValidation(hostname, tokens[1])

	if err != nil {
		tools.SendErrorResponse(w, option.INVALID_AUTHORIZATION, http.StatusUnauthorized)
		return
	}

	ctx := context.WithValue(r.Context(), option.CONTEXT_PLATFORM, platform)

	next.ServeHTTP(w, r.WithContext(ctx))
}
