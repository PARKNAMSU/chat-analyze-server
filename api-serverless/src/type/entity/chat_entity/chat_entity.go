package chat_entity

import "time"

type ChatEntity struct {
	Id          *int       `json:"id" db:"id"`
	Domain      *string    `json:"domain" db:"domain"`
	CreatedUser *string    `json:"createdUser" db:"created_user"`
	Type        *int       `json:"type" db:"type"`
	Password    *int       `json:"password" db:"password"`
	MaxUser     *int       `json:"maxUser" db:"max_user"`
	CreatedAt   *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   *time.Time `json:"updatedAt" db:"updated_at"`
}
