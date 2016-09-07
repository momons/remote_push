package request

// トークン登録リクエストEntity.
type RegisterToken struct {
	// 通知コード.
	NotificationCode string `json:"notificationCode"`
	// トークン.
	Token string `json:"token"`
	// カスタムパラメータ.
	CustomParams string `json:"custom_params"`
}

func (entity *RegisterToken) Convert(
	object map[string]interface{},
) {
	if value, ok := object["notificationCode"].(string); ok {
		entity.NotificationCode = value
	}
	if value, ok := object["token"].(string); ok {
		entity.Token = value
	}
	if value, ok := object["custom_params"].(string); ok {
		entity.CustomParams = value
	}
}
