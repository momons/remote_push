package response

// 共通レスポンスEntity.
type Common struct {
	// ステータス.
	Status Status `json:"status"`
	// メタ.
	Meta interface{} `json:"meta"`
}
