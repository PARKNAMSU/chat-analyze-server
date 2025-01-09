package models

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type SocketRouter func(http.ResponseWriter, *http.Request, *websocket.Conn)
