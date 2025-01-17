package platform_entity

import "time"

type PartnerPlatformEntity struct {
	Id          int        `db:"id"`
	Domain      *string    `db:"domain"`        // 사이트 도메인
	IpAddr      *string    `db:"ip_addr"`       // 플랫폼 등록 시 Ip 주소
	UserKeyType *string    `db:"user_key_type"` // 사이트 유저 고유키 데이터타입
	Name        *string    `db:"name"`          // 사이트 이름
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}

type PartnerPlatformAPIKeyEntity struct {
	PlatformId  *int       `db:"platform_id"`
	ApiHashKey  *string    `db:"api_hash_key"` // 암호화된 API KEY
	AvaliableAt *time.Time `db:"avaliable_at"` // 사용 가능 시간
	ExpiredAt   *time.Time `db:"expired_at"`   // 만료 시간
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}

type PartnerPlatformPermissionEntity struct {
	PlatformId    *int       `db:"platform_id"`
	Type          *int       `db:"type"`           // 1: PREMIUM, 2: NORMAL
	PermissionBit *int       `db:"permission_bit"` // 권한 비트
	MaxChatNum    *int       `db:"max_chat_num"`   // 최대 채팅 수
	MaxAccessNum  *int       `db:"max_access_num"` // 최대 접속 수
	MaxUserNum    *int       `db:"max_user_num"`   // 최대 유저 수
	CreatedAt     *time.Time `db:"created_at"`
	UpdatedAt     *time.Time `db:"updated_at"`
}

type PartnerPlatformConfigEntity struct {
	PlatformId       *int       `db:"platform_id"`
	IsExitUserDelete *int       `db:"is_exit_user_delete"` // 유저가 모든 채팅방에서 나갈 때 데이터 삭제 여부
	CreatedAt        *time.Time `db:"created_at"`
	UpdatedAt        *time.Time `db:"updated_at"`
}

type PartnerPlatformUserEntity struct {
	PlatformId      *int       `db:"platform_id"`
	UserKey         *string    `db:"user_key"`
	UserInformation *string    `db:"user_information"`
	UpdatedAt       *time.Time `db:"updated_at"`
	CreatedAt       *time.Time `db:"created_at"`
}
