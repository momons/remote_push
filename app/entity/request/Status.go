package request

import (
	"../../constants"
)

// ステータスリクエストEntity.
type Status struct {
	// アプリトークン.
	AppToken string `json:"appToken"`
	// アプリバージョン.
	AppVersion string `json:"appVersion"`
	// プラットフォーム.
	Platform string `json:"platform"`
	// ユーザコード.
	UserCode string `json:"userCode"`
	// ユーザ名.
	UserName string `json:"userName"`
	// アクセストークン.
	AccessToken string `json:"accessToken"`
}

// アプリトークンチェック.
func (status *Status) IsValidAppToken() bool {
	return constants.AppToken == status.AppToken
}
