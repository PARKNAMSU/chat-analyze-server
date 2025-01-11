package chat_entity

import "time"

/*
채팅방 INDEX 스키마
Description:

	PK: ChatId
*/
type ChatEntity struct {
	ChatId      *int       `db:"chatId"`      // 채팅방 아이디 - not null, auto_increment, primary key
	IpAddr      *string    `db:"ipAddr"`      // 생성된 시점 아이피 주소 - not null
	UserId      *int       `db:"userId"`      // 채팅방 생성자 아이디 - not null
	Status      *int       `db:"status"`      // 채팅방 상태 0:삭제됨, 1:정상, 2:정지됨 - not null, default: 1
	ForbiddenAt *time.Time `db:"forbiddenAt"` // 채팅방 정지 시간 - nullable
	UpdatedAt   *time.Time `db:"updatedAt"`   // 채팅방 수정 시간 - nullable
	CreatedAt   *time.Time `db:"createdAt"`   // 채팅방 생성 시간 - not null, default: now()
}

/*
채팅방 정보 스키마
Description:

	PK: ChatId
	Relation: ChatEntity.ChatId = ChatInformationEntity.ChatId (1:1)
*/
type ChatInformationEntity struct {
	ChatId     *int       `db:"chatId"`     // ChatEntity.ChatId - not null, PK
	ChatName   *string    `db:"chatName"`   // 채팅방 이름 - nullable
	ChatType   *int       `db:"chatType"`   // 채팅방 타입 0:개인, 1:공개그룹, 2:비공개그룹 - not null, default: 0
	Password   *int       `db:"password"`   // 채팅방 비밀번호 - nullable
	IsMaxLimit *int       `db:"isMaxLimit"` // 채팅방 최대 인원 제한 여부 0:제한없음, 1:제한있음 - not null, default: 0
	MaxUser    *int       `db:"maxUser"`    // 채팅방 최대 인원 - nullable
	UpdatedAt  *time.Time `db:"updatedAt"`  // 채팅방 수정 시간 - nullable
	CreatedAt  *time.Time `db:"createdAt"`  // 채팅방 생성 시간 - not null, default: now()
}

/*
채팅방 참가 유저 스키마
Description:

	PK: ChatId, UserId
	Relation: ChatEntity.ChatId = ChatUserEntity.ChatId (1:N), UserEntity.UserId = ChatUserEntity.UserId (1:N)
*/
type ChatUserEntity struct {
	ChatId      *string    `db:"chatId"`      // ChatEntity.ChatId - not null, PK
	UserId      *string    `db:"userId"`      // ChatUserEntity.UserId - not null, PK
	UserType    *int       `db:"userType"`    // ChatUserEntity.UserType 1:생성자, 2:참가자 - not null, default: 2
	Status      *int       `db:"status"`      // 채팅방 상태 0:영구정지, 1:정상, 2:정지됨 - not null, default: 1
	ForbiddenAt *time.Time `db:"forbiddenAt"` // 정지 시간 - nullable
	AttendedAt  *time.Time `db:"attendedAt"`  // 참가 시간 - not null, default: now()
}
