package manager

import (
	"../entity/database"
	"log"
	"time"
)

// 通知ユーザデータ用マネージャ.
type NotificationUsers struct {
}

// インスタンス.
var instanceNotificationUsers *NotificationUsers

// インスタンス取得.
func GetNotificationUsers() *NotificationUsers {
	if instanceNotificationUsers == nil {
		instanceNotificationUsers = &NotificationUsers{}
	}
	return instanceNotificationUsers
}

// 通知コード、トークンより取得
func (manager *NotificationUsers) Select(
	notificationCode string,
	notificationToken string,
) *database.NotificationUsers {

	var entity database.NotificationUsers

	err := Db.Where(
		"notification_code = ? AND notification_token = ?",
		notificationCode,
		notificationToken,
	).First(
		&entity,
	).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &entity
}

// 削除＆作成.
func (manager *NotificationUsers) UpdateInsert(
	notificationCode string,
	notificationToken string,
	platform string,
	customParams string,
) bool {

	Db.Begin()

	nowAt := time.Now()

	// 既存にある場合削除
	entity := manager.Select(notificationCode, notificationToken)
	if entity != nil {
		entity.Platform = platform
		entity.CustomParams = customParams
		entity.UpdateAt = nowAt
		err := Db.Update(entity).Error
		if err != nil {
			log.Println(err)
			Db.Rollback()
			return false
		}
	} else {
		insertEntity := database.NotificationUsers{
			NotificationCode:  notificationCode,
			NotificationToken: notificationToken,
			Platform:          platform,
			CustomParams:      customParams,
			UpdateAt:          nowAt,
			CreateAt:          nowAt,
		}
		err := Db.Create(&insertEntity).Error
		if err != nil {
			log.Println(err)
			Db.Rollback()
			return false
		}
	}

	Db.Commit()

	return true
}
