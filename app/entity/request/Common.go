package request

// 共通リクエストEntity.
type Common struct {
	// ステータス.
	Status Status `json:"status"`
	// パラメータ.
	Params map[string]interface{} `json:"params"`
}
