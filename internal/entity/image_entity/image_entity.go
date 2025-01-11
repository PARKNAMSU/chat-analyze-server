package image_entity

import "time"

/*
기본 이미지 스키마
*/
type ImageSchema struct {
	OriginFile   *string    `db:"originFile"`
	JpegFile     *string    `db:"jpegFile"`
	WebpFile     *string    `db:"webpFile"`
	JpegSize     *int       `db:"jpegSize"`
	WebpSize     *int       `db:"webpSize"`
	OriginBucket *string    `db:"originBucket"`
	JpegBucket   *string    `db:"jpegBucket"`
	WebpBucket   *string    `db:"webpBucket"`
	CreatedAt    *time.Time `db:"createdAt"` // 생성 시간 - not null, default: now()
}

/*
유저 프로필 이미지 스키마
Description:

	ImageSchema 상속
	PK: ImageId
	INDEX: (userId)
*/
type UserTitleImageSchema struct {
	ImageId *string `db:"imageId"` // 이미지 생성 아이디 - not null
	UserId  *int    `db:"userId"`  // 유저 아이디 - not null, PK
	Device  *int    `db:"device"`  // 이미지 디바이스 0:모바일, 1:웹 - not null, default: 0
	ImageSchema
}
