package user_entity

import (
	"time"

	"chat-platform-api.com/chat-platform-api/src/type/entity/common_entity"
)

type UserEntity struct {
	Id          *int       `json:"id" db:"id"`
	Status      *int       `json:"status" db:"status"`
	IpAddr      *string    `json:"ipAddr" db:"ip_addr"`
	LastLoginAt *time.Time `json:"lastLoginAt" db:"last_login_at"`
	common_entity.TableTimestamp
}

type UserInformationEntity struct {
	UserId         *int    `json:"userId" db:"user_id"`
	Email          *string `json:"email" db:"email"`
	Password       *string `json:"password" db:"password"`
	Name           *string `json:"name" db:"names"`
	Authentication *int    `json:"authentication" db:"authentication"`
	AuthType       *string `json:"authType" db:"auth_type"`
	common_entity.TableTimestamp
}

type UserOauthEntity struct {
	UserId    *int    `json:"userId" db:"user_id"`
	OauthId   *string `json:"oauthId" db:"oauth_id"`
	OauthHost *string `json:"oauthHost" db:"oauth_host"`
	common_entity.TableTimestamp
}

// revoke token 검증을 위한 테이블
type UserRefreshTokenEntity struct {
	UserId   *int    `json:"userId" db:"user_id"`
	Token    *string `json:"token" db:"token"`
	IpAddr   *string `json:"ipAddr" db:"ip_addr"`
	DeviceId *string `json:"deviceId" db:"device_id"`
	common_entity.TableTimestamp
}
