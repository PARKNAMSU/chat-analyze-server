package common_dto

type DefaultRequest struct {
	Router string `json:"router"` // 호출한 작업의 카테고리 - required
}
