package polling

import (
	"../../entity/database"
	"../../manager"
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
		if messages != nil {

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

// メッセージ送信
func (service *PollingService) sendMessage(
	message database.NotificationMessages,
	users *[]database.NotificationUsers,
) {

	for _, user := range *users {

		// プラットフォーム指定あり
		if (message.IsiOS() && !user.IsiOS()) || (message.IsAndroid() && !user.IsAndroid()) {
			continue
		}

	}

}
