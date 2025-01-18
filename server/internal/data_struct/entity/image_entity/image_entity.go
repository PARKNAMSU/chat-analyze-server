package image_entity

import "time"

/*
기본 이미지 엔티티
*/
type ImageEntity struct {
	OriginFile   *string    `db:"originFile"`   // 원본 이미지 파일 경로
	JpegFile     *string    `db:"jpegFile"`     // jpeg 이미지 파일 경로
	WebpFile     *string    `db:"webpFile"`     // webp 이미지 파일 경로
	JpegSize     *int       `db:"jpegSize"`     // jpeg 이미지 파일 크기
	WebpSize     *int       `db:"webpSize"`     // webp 이미지 파일 크기
	OriginBucket *string    `db:"originBucket"` // 원본 이미지 저장 버킷
	JpegBucket   *string    `db:"jpegBucket"`   // jpeg 이미지 저장 버킷
	WebpBucket   *string    `db:"webpBucket"`   // webp 이미지 저장 버킷
	CreatedAt    *time.Time `db:"createdAt"`    // 생성 시간 - not null, default: now()
}

/*
유저 프로필 이미지 엔티티
*/
type UserProfileImageEntity struct {
	ImageId *string `db:"imageId"` // 이미지 생성 아이디 - not null, PK
	UserId  *int    `db:"userId"`  // 유저 아이디 - not null
	ImageEntity
}

/*
메세지 이미지 파일
*/
type MessageImageEntity struct {
	ImageId *string `db:"imageId"` // 이미지 생성 아이디 - not null, PK
	UserId  *int    `db:"userId"`  // 유저 아이디 - not null,
	ChatId  *int    `db:"chatId"`  // 채팅방 아이디 - not null,
	ImageEntity
}
