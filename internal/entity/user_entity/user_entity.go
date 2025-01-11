package user_entity

import "time"

/*
사용자 INDEX 스키마
Description:

	PK: UserId
*/
type UserSchema struct {
	UserId     *int       `db:"userId"`     // 사용자 아이디 - not null, auto_increment, primary key
	Status     *int       `db:"status"`     // 사용자 상태 0:탈퇴, 1:정상, 2:정지 - not null, default: 1
	IpAddr     *string    `db:"ipAddr"`     // 사용자 아이피 주소 - not null
	Auth       *int       `db:"auth"`       // 사용지 인증 여부 0:미인증, 1:인증 - not null, default: 0
	AuthMethod *int       `db:"authMethod"` // 사용자 인증 방법 0:이메일, 1:모바일폰 - nullable, default: 0
	DeviceId   *string    `db:"deviceId"`   // 사용자 디바이스 아이디 - nullable
	Locale     *string    `db:"locale"`     // 사용자 지역 - not null, default: "XX"
	UpdatedAt  *time.Time `db:"updatedAt"`  // 유저 수정 시간 - nullable
	CreatedAt  *time.Time `db:"createdAt"`  // 유저 생성 시간 - not null, default: now()
}

/*
사용자 정보 스키마
Description:

	PK: UserId
	relation: UserSchema.UserId = UserInformationSchema.UserId (1:1)
*/
type UserInformationSchema struct {
	UserId    *int       `db:"userId"`    // 사용자 아이디 - not null, PK
	Name      *int       `db:"name"`      // 사용자 이름 - nullable
	Email     *int       `db:"email"`     // 사용자 로그인 이메일 - not null
	Password  *int       `db:"password"`  // 사용자 로그인 비밀번호 - not null
	Gender    *int       `db:"gender"`    // 사용자 성별 0:남성, 1:여성 - not null default: 0
	UpdatedAt *time.Time `db:"updatedAt"` // 유저정보 수정 시간 - nullable
	CreatedAt *time.Time `db:"createdAt"` // 유저정보 생성 시간 - not null, default: now()
}

/*
사용자 oauth 등록 스키마
Description:

	PK: UserId
	relation: UserSchema.UserId = UserOauthSchema.UserId (1:1)
*/
type UserOauthSchema struct {
	UserId    *int       `db:"userId"`    // 사용자 아이디 - not null, PK
	OauthKey  *string    `db:"oauthKey"`  // oauth 키 - not null, PK
	OauthType *string    `db:"oauthType"` // 사용 oauth 종류 ex) google.com - not null
	UpdatedAt *time.Time `db:"updatedAt"` // 수정 시간 - nullable
	CreatedAt *time.Time `db:"createdAt"` // 생성 시간 - not null, default: now()
}
