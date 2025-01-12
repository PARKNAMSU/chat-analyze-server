package common_middleware

import (
	"net/http"
	"os"

	"chat-analyze.com/chat-analyze-server/internal/options"
	"chat-analyze.com/chat-analyze-server/internal/tools"
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
		tools.SendErrorResponse(w, options.INVALID_API_KEY, http.StatusUnauthorized)
		return
	}
	next(w, r)
}
