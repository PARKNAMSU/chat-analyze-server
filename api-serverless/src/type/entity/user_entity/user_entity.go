package user_entity

import "time"

type UserEntity struct {
	Id        *int       `json:"id" db:"id"`
	Status    *int       `json:"status" db:"status"`
	IpAddr    *string    `json:"ipAddr" db:"ip_addr"`
	CreatedAt *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}

type UserInformationEntity struct {
	UserId         *int       `json:"userId" db:"user_id"`
	Email          *string    `json:"email" db:"email"`
	Name           *string    `json:"name" db:"names"`
	Authentication *int       `json:"authentication" db:"authentication"`
	AuthType       *string    `json:"authType" db:"auth_type"`
	CreatedAt      *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt      *time.Time `json:"updatedAt" db:"updated_at"`
}

type UserOauth struct {
	UserId    *int       `json:"userId" db:"user_id"`
	OauthId   *string    `json:"oauthId" db:"oauth_id"`
	OauthHost *string    `json:"oauthHost" db:"oauth_host"`
	CreatedAt *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}

// revoke token 검증을 위한 테이블
type UserRefreshToken struct {
	UserId    *int       `json:"userId" db:"user_id"`
	Token     *string    `json:"token" db:"token"`
	IpAddr    *string    `json:"ipAddr" db:"ip_addr"`
	DeviceId  *string    `json:"deviceId" db:"device_id"`
	CreatedAt *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}
