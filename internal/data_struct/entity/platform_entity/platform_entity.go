package platform_entity

import "time"

type PartnerPlatform struct {
	Id          int        `db:"id"`
	Domain      *string    `db:"domain"`
	IpAddr      *string    `db:"ip_addr"`
	UserKeyType *string    `db:"user_key_type"`
	Name        *string    `db:"name"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}

type PartnerPlatformAPIKey struct {
	PlatformId      *int       `db:"platform_id"`
	ApiKey          *string    `db:"api_key"`
	AuthBit         *int       `db:"auth_bit"`
	MaxSocketAccess *int       `db:"max_socket_access"`
	MaxChatNumber   *int       `db:"max_chat_number"`
	AvaliableAt     *time.Time `db:"avaliable_at"`
	ExpiredAt       *time.Time `db:"expired_at"`
	CreatedAt       *time.Time `db:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at"`
}

type PartnerPlatformUser struct {
	PlatformId      *int       `db:"platform_id"`
	UserKey         *string    `db:"user_key"`
	UserInformation *string    `db:"user_information"`
	UpdatedAt       *time.Time `db:"updated_at"`
	CreatedAt       *time.Time `db:"created_at"`
}
