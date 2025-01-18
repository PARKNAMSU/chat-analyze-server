package index_middleware

import (
	"net/http"

	"chat-analyze.com/chat-analyze-server/internal/data_struct/model/common_model"
	"chat-analyze.com/chat-analyze-server/internal/tools"
)

// 미들웨어 체이닝에 적용할 미들웨어 생성 함수
func Middleware(handler common_model.Middleware, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, next)
	})
}

// 라우터 실행 전 미들웨어들을 체이닝하여 실행하는 함수
func MiddlewareChaining(router http.HandlerFunc, middlewares ...common_model.Middleware) http.HandlerFunc {
	slice := tools.CustomSlice[common_model.Middleware]{
		Slice: middlewares,
	}

	for _, middleware := range slice.Reverse() {
		router = Middleware(middleware, router)
	}
	return router
}
