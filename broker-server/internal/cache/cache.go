package cache

import (
	"chat-analyze.com/chat-analyze-server/internal/data_struct/model/common_model"
	"github.com/gorilla/websocket"
)

var (
	// 추후 해당 캐시를 redis 서버로 이전
	connCache = make(map[int]map[int]*websocket.Conn)
)

func SetChatCache(conn *common_model.GetConnectData) {
	_, isOk := connCache[conn.ChatId]
	if !isOk {
		connCache[conn.ChatId] = make(map[int]*websocket.Conn)
	}
	connCache[conn.ChatId][conn.UserId] = conn.Conn
}

func DeleteChatCache(chatId int, userId int) {
	if _, isOk := connCache[chatId]; !isOk {
		return
	}
	delete(connCache[chatId], userId)
}

func ClearChatCache(chatId int) {
	if _, isOk := connCache[chatId]; !isOk {
		return
	}
	delete(connCache, chatId)
}

func GetUserConn(userId int, chatId int) *websocket.Conn {
	chatData, findChat := connCache[chatId]
	if !findChat {
		return nil
	}
	userConn, findUser := chatData[userId]
	if !findUser {
		return nil
	}
	return userConn
}

func GetChatConns(chatId int) map[int]*websocket.Conn {
	chatData, findChat := connCache[chatId]
	if !findChat {
		return nil
	}
	return chatData
}
