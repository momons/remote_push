package response

// ステータスレスポンスEntity.
type Status struct {
	// コード.
	Code int `json:"code"`
	// メッセージ.
	Message string `json:"message"`
}
