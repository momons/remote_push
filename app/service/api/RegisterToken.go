package api

import (
	"../../constants"
	"../../entity/request"
	"../../entity/response"
	"../../manager"
	"../../util"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

// トークン登録サービス.
type RegisterToken struct {
	// 通知マスタマネージャ.
	notificationsManager *manager.Notifications
	// 通知ユーザデータマネージャ.
	notificationUsersManager *manager.NotificationUsers
}

// インスタンス.
var instanceRegisterToken *RegisterToken

// インスタンス取得.
func GetRegisterToken() *RegisterToken {
	if instanceRegisterToken == nil {
		instanceRegisterToken = &RegisterToken{
			notificationsManager:     manager.GetNotifications(),
			notificationUsersManager: manager.GetNotificationUsers(),
		}
	}
	return instanceRegisterToken
}

// 受信.
func (service *RegisterToken) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckAppToken(w, req)
	if !isSuccess {
		return
	}
	params := request.RegisterToken{}
	params.Convert(requestEntity.Params)

	// 通知コードチェック
	hasNotificationCode := service.notificationsManager.HasNotificationCode(params.NotificationCode)
	if !hasNotificationCode {
		w.WriteJson(util.ErrorResponseEntity(http.StatusInternalServerError, constants.MessageE0010))
		return
	}

	// 追加
	isSuccess = service.notificationUsersManager.UpdateInsert(
		params.NotificationCode,
		params.Token,
		requestEntity.Status.Platform,
		params.CustomParams,
	)
	if !isSuccess {
		w.WriteJson(util.ErrorResponseEntity(http.StatusInternalServerError, constants.MessageE4001))
		return
	}

	// レスポンス作成
	metaEntity := response.RegisterToken{}
	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}
