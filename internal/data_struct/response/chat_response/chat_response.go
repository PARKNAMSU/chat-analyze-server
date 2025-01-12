package chat_response

type SendTextResponse struct {
	MessageId int    `json:"messageId"` // 메시지 아이디 - required
	Name      string `json:"name"`      // 사용자 이름 - required
	Message   string `json:"message"`   // 메시지 내용 - required
	Status    int    `json:"status"`    // 메시지 상태 0:삭제됨, 1:정상 - required
	CreatedAt string `json:"createdAt"` // 메시지 생성 시간 Format: "YYYY-MM-DD HH:mm:ss" - required
}
