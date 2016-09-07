package request

// メッセージ登録リクエストEntity.
type RegisterMessage struct {
	// 通知コード.
	NotificationCode string `json:"notificationCode"`
	// 送信日時
	SendAt string `json:"sendAt"`
	// メッセージ.
	Message string `json:"message"`
	// 送信条件 .
	SendCondition string `json:"sendCondition"`
}

func (entity *RegisterMessage) Convert(
object map[string]interface{},
) {
	if value, ok := object["notificationCode"].(string); ok {
		entity.NotificationCode = value
	}
	if value, ok := object["sendAt"].(string); ok {
		entity.SendAt = value
	}
	if value, ok := object["message"].(string); ok {
		entity.Message = value
	}
	if value, ok := object["sendCondition"].(string); ok {
		entity.SendCondition = value
	}
}
