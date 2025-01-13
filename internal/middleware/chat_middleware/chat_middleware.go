package chat_middleware

import (
	"context"
	"fmt"
	"net/http"

	"chat-analyze.com/chat-analyze-server/internal/option"
	"chat-analyze.com/chat-analyze-server/internal/tools"
)

func AttendChatMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	userId, chatId, err := tools.AttendRoom(r.Header)

	if err != nil || chatId == 0 || userId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Not Exist Chat"))
	}

	tools.PrintInfoLog("AttendChatMiddleware", fmt.Sprintf("Client connected to group: %d\n", chatId))

	ctx := context.WithValue(r.Context(), option.CONTEXT_USER_ID, userId)
	ctx = context.WithValue(ctx, option.CONTEXT_CHAT_ID, chatId)

	next.ServeHTTP(w, r.WithContext(ctx))
}
