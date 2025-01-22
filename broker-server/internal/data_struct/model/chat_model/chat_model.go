package chat_model

import "time"

type MessageData struct {
	ChatId    int       `json:"chatId"`
	UserId    int       `json:"userId"`
	Text      string    `json:"text"`
	Order     int       `json:"order"`
	CreatedAt time.Time `json:"created_at"`
}
