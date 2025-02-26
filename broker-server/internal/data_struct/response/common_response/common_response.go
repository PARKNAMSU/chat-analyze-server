package common_response

// 클라이언트에 메세지 전송 시 기본적으로 사용되는 구조체
type ResponseDefault struct {
	Message *string `json:"message,omitempty"` // 메시지 - optional
	Status  int     `json:"status"`            // 상태 - required
}
