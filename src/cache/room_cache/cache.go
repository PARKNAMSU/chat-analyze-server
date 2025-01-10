package room_cache

var (
	// 추후 해당 캐시를 redis 서버로 이전
	userRooms = make(map[int]map[int]bool)
)

func SetChatCache(chatId int, userId int) {
	_, isOk := userRooms[chatId]
	if !isOk {
		userRooms[chatId] = make(map[int]bool)
	}
	userRooms[chatId][userId] = true
}

func DeleteChatCache(chatId int, userId int) {
	if _, isOk := userRooms[chatId]; !isOk {
		return
	}
	delete(userRooms[chatId], userId)
}

func ClearChatCache(chatId int) {
	if _, isOk := userRooms[chatId]; !isOk {
		return
	}
	delete(userRooms, chatId)
}
