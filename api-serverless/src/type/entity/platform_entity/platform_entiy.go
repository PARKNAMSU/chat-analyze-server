package platform_entity

import "time"

type PlatformEntity struct {
	Domain              *string    `json:"domain" db:"domain"`
	UserId              *int       `json:"userId" db:"user_id"`
	Status              *int       `json:"status" db:"status"`
	Name                string     `json:"name" db:"name"`
	PlatformUserKeyType *string    `json:"platformUserKeyType" db:"platform_user_key_type"`
	CreatedAt           *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt           *time.Time `json:"updatedAt" db:"updated_at"`
}

type PlatformPermissionEntity struct {
	Domain        *string    `json:"domain" db:"domain"`
	PermissionBit *int       `json:"permissionBit" db:"permission_bit"`
	Grade         *string    `json:"grade" db:"grade"`
	MaxChatNum    *int       `json:"maxChatNum" db:"max_chat_num"`
	MaxAccessNum  *int       `json:"maxAccessNum" db:"max_access_num"`
	MaxUserNum    *int       `json:"maxUserNum" db:"max_user_num"`
	CreatedAt     *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt     *time.Time `json:"updatedAt" db:"updated_at"`
}

type PlatformAccessKeyEntity struct {
	Domain    *string    `json:"domain" db:"domain"`
	AccessKey *string    `json:"accessKey" db:"access_key"`
	SecretKey *string    `json:"secretKey" db:"secret_key"`
	ExpiredAt *time.Time `json:"expiredAt" db:"expired_at"`
	CreatedAt *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}
