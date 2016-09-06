package request

// トークン登録リクエストEntity.
type RegisterToken struct {
	// トークン.
	Token string `json:"token"`
}

func (entity *RegisterToken) Convert(
	object map[string]interface{},
) {
	if value, ok := object["token"].(string); ok {
		entity.Token = value
	}
}
