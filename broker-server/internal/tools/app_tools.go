package tools

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"chat-analyze.com/chat-analyze-server/internal/data_struct/model/common_model"
	"github.com/gorilla/websocket"
)

var (
	websocketCreator = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// 모든 요청 허용
			return true
		},
	}
)

func GetWebSocket(w http.ResponseWriter, r *http.Request) *websocket.Conn {
	conn, err := websocketCreator.Upgrade(w, r, nil)

	if err != nil {
		PrintErrorLog("getWebSocket", err.Error())
		return nil
	}
	log.Println("Client connected")
	return conn
}

func GetConnData(w http.ResponseWriter, r *http.Request) (
	connData *common_model.GetConnectData,
	err error,
) {
	headerUserId := r.Header.Get("userId")
	headerChatId := r.Header.Get("chatId")

	if headerUserId == "" || headerChatId == "" {
		return nil, errors.New("Not Exist UserId or ChatId")
	}

	userId, err := strconv.Atoi(headerUserId)
	chatId, err := strconv.Atoi(headerChatId)

	if err != nil {
		return nil, err
	}

	return &common_model.GetConnectData{
		UserId: userId,
		ChatId: chatId,
		Conn:   GetWebSocket(w, r),
	}, nil
}
