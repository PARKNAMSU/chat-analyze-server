package tools

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"chat-analyze.com/chat-analyze-server/src/cache/room_cache"
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

func AttendRoom(header http.Header) (
	userId int,
	chatId int,
	err error,
) {
	userId = 0
	chatId = 0

	headerUserId := header.Get("userId")
	headerChatId := header.Get("chatId")

	if headerUserId == "" || headerChatId == "" {
		return userId, chatId, errors.New("Not Exist UserId or ChatId")
	}

	userId, err = strconv.Atoi(headerUserId)
	chatId, err = strconv.Atoi(headerChatId)

	if err != nil {
		return userId, chatId, err
	}

	room_cache.SetChatCache(chatId, userId)

	return userId, chatId, nil
}
