package cache

import "github.com/gorilla/websocket"

var (
	USER_GROUP = make(map[int]map[*websocket.Conn]bool)
)
