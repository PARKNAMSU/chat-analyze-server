package common_model

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Middleware func(http.ResponseWriter, *http.Request, http.HandlerFunc)

type GetConnectData struct {
	Conn   *websocket.Conn // 웹소켓 연결 객체 - required
	UserId int             // 사용자 아이디 - required
	ChatId int             // 참가하고 있는 채팅방 아이디 - required
}

type ContextKey string
