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

// メッセージ登録サービス.
type RegisterMessage struct {
	// 通知マスタマネージャ.
	notificationsManager *manager.Notifications
	// 通知メッセージデータマネージャ.
	notificationMessagesManager *manager.NotificationMessages
}

// インスタンス.
var instanceRegisterMessage *RegisterMessage

// インスタンス取得.
func GetRegisterMessage() *RegisterMessage {
	if instanceRegisterMessage == nil {
		instanceRegisterMessage = &RegisterMessage{
			notificationsManager:        manager.GetNotifications(),
			notificationMessagesManager: manager.GetNotificationMessages(),
		}
	}
	return instanceRegisterMessage
}

// 受信.
func (service *RegisterMessage) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckAppToken(w, req)
	if !isSuccess {
		return
	}
	params := request.RegisterMessage{}
	params.Convert(requestEntity.Params)

	// 通知コードチェック
	hasNotificationCode := service.notificationsManager.HasNotificationCode(params.NotificationCode)
	if !hasNotificationCode {
		w.WriteJson(util.ErrorResponseEntity(http.StatusInternalServerError, constants.MessageE0010))
		return
	}

	// 追加
	isSuccess, messageCode := service.notificationMessagesManager.UpdateInsert(
		nil,
		params.NotificationCode,
		params.SendAt,
		requestEntity.Status.Platform,
		params.Message,
		params.SendCondition,
	)
	if !isSuccess {
		w.WriteJson(util.ErrorResponseEntity(http.StatusInternalServerError, constants.MessageE4002))
		return
	}

	// レスポンス作成
	metaEntity := response.RegisterMessage{
		MessageCode: *messageCode,
	}
	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}
