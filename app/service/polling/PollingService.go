package polling

import (
	"../../entity/database"
	"../../manager"
	"encoding/json"
	"fmt"
	"github.com/alexjlockwood/gcm"
	"github.com/anachronistic/apns"
	"strings"
	"time"
)

// ポーリングサービス.
type PollingService struct {
	// 通知マネージャ.
	notificationsManager *manager.Notifications
	// 通知ユーザマネージャ.
	notificationUsersManager *manager.NotificationUsers
	// 通知メッセージマネージャ.
	notificationMessagesManager *manager.NotificationMessages
}

// インスタンス.
var instancePollingService *PollingService

// インスタンス取得.
func GetPollingService() *PollingService {
	if instancePollingService == nil {
		instancePollingService = &PollingService{
			notificationsManager:        manager.GetNotifications(),
			notificationUsersManager:    manager.GetNotificationUsers(),
			notificationMessagesManager: manager.GetNotificationMessages(),
		}
	}
	return instancePollingService
}

// ポーリングスタート
func (service *PollingService) Start() {
	for {

		// 送信対象メッセージ取得
		messages := service.notificationMessagesManager.SelectUnSent()
		if messages != nil && len(*messages) > 0 {

			// 送信対象者取得
			users := service.notificationUsersManager.SelectAllUsers()
			for _, message := range *messages {
				if users != nil {
					// メッセージ送信
					go service.sendMessage(message, users)
				}

				// 送信済みに更新
				service.notificationMessagesManager.UpdateSendFlag(message.MessageCode)

				// スリープ 1秒
				time.Sleep(1 * time.Second)
			}
		}

		// スリープ 5秒
		time.Sleep(5 * time.Second)
	}
}

// メッセージ送信.
func (service *PollingService) sendMessage(
	message database.NotificationMessages,
	users *[]database.NotificationUsers,
) {

	// 通知マスタ取得
	notification := service.notificationsManager.Select(message.NotificationCode)

	// iOSペイロード作成
	iOSPayload, addInfo := service.createiOSPayload(message.Message)

	// Androidトークン
	androidTokens := []string{}

	for _, user := range *users {

		// プラットフォーム指定あり
		if (message.IsiOS() && !user.IsiOS()) || (message.IsAndroid() && !user.IsAndroid()) {
			continue
		}

		if user.IsiOS() {
			service.sendMessageiOS(
				user.NotificationToken,
				iOSPayload,
				addInfo,
				notification.CertFileName,
				notification.KeyFileName,
			)
		} else if user.IsAndroid() {
			// トークンを追加していく
			androidTokens = append(androidTokens, user.NotificationToken)
		}
	}

	// 最後にAndroidが複数あったらメッセージ送信
	if len(androidTokens) > 0 {
		// メッセージ作成
		message := service.createAndroidPayload(message.Message)
		// 送信
		service.sendMessageAndroid(androidTokens, message, notification.ApiKey)
	}
}

// メッセージよりiOSペイロードを作成する.
func (service *PollingService) createiOSPayload(
	message string,
) (*apns.Payload, string) {

	payload := apns.NewPayload()
	addInfo := ""

	// JSON変換
	var jsonData map[string]interface{}
	dec := json.NewDecoder(strings.NewReader(message))
	dec.Decode(&jsonData)

	if value, ok := jsonData["alert"].(string); ok {
		payload.Alert = value
	}
	if value, ok := jsonData["badge"].(int); ok {
		payload.Badge = value
	}
	if value, ok := jsonData["sound"].(string); ok {
		payload.Sound = value
	}
	if value, ok := jsonData["contentAvailable"].(int); ok {
		payload.ContentAvailable = value
	}
	if value, ok := jsonData["category"].(string); ok {
		payload.Category = value
	}
	if value, ok := jsonData["addInfo"].(string); ok {
		addInfo = value
	}

	return payload, addInfo
}

// iOSメッセージ送信.
func (service *PollingService) sendMessageiOS(
	deviceToken string,
	payload *apns.Payload,
	addInfo string,
	certFilePath string,
	keyFilePath string,
) {

	pn := apns.NewPushNotification()
	pn.DeviceToken = deviceToken
	pn.Set("addInfo", addInfo)
	pn.AddPayload(payload)

	client := apns.NewClient("gateway.sandbox.push.apple.com:2195", certFilePath, keyFilePath)
	resp := client.Send(pn)

	alert, _ := pn.PayloadString()

	fmt.Println("Alert\t", alert)
	fmt.Println("Succe\t", resp.Success)
	fmt.Println("Error\t", resp.Error)
}

// メッセージよりAndroidペイロードを作成する.
func (service *PollingService) createAndroidPayload(
	message string,
) map[string]interface{} {

	// JSON変換
	var jsonData map[string]interface{}
	dec := json.NewDecoder(strings.NewReader(message))
	dec.Decode(&jsonData)

	return jsonData
}

// Androidメッセージ送信.
func (service *PollingService) sendMessageAndroid(
	deviceTokens []string,
	message map[string]interface{},
	apiKey string,
) {
	// メッセージ作成
	msg := gcm.NewMessage(message, deviceTokens...)

	// 送信
	sender := &gcm.Sender{ApiKey: apiKey}
	response, err := sender.Send(msg, 2)
	if err != nil {
		fmt.Println("Failed to send message:", err)
		return
	}

	fmt.Println("Alert\t", response)
}
