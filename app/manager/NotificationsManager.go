package manager

import (
	"../entity/database"
	"log"
	"github.com/jinzhu/gorm"
)

// 通知マスタ用マネージャ.
type Notifications struct {
}

// インスタンス.
var instanceNotifications *Notifications

// インスタンス取得.
func GetNotifications() *Notifications {
	if instanceNotifications == nil {
		instanceNotifications = &Notifications{}
	}
	return instanceNotifications
}

// 通知コードより取得
func (manager *Notifications) Select(
	notificationCode string,
) *database.Notifications {

	var entity database.Notifications

	err := Db.Where(
		"notification_code = ?",
		notificationCode,
	).First(
		&entity,
	).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Println(err)
		}
		return nil
	}

	return &entity
}

// 通知コードが存在するかチェック
func (manager *Notifications) HasNotificationCode(
	notificationCode string,
) bool {
	entity := manager.Select(notificationCode)
	if entity == nil {
		return false
	}
	return true
}
