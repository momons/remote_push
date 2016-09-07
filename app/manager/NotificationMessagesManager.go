package manager

import (
	"../constants"
	"../entity/database"
	"../util"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

// 通知メッセージデータ用マネージャ.
type NotificationMessages struct {
}

// インスタンス.
var instanceNotificationMessages *NotificationMessages

// インスタンス取得.
func GetNotificationMessages() *NotificationMessages {
	if instanceNotificationMessages == nil {
		instanceNotificationMessages = &NotificationMessages{}
	}
	return instanceNotificationMessages
}

// 未送信メッセージ取得
func (manager *NotificationMessages) SelectUnSent() *[]database.NotificationMessages {

	var entities []database.NotificationMessages

	// 現在時刻を変換
	nowAt := util.StringFromDate(time.Now())

	err := Db.Where(
		"(send_at <= ? OR send_at = '') AND is_send = ?",
		nowAt,
		constants.SendTypeNone,
	).Order(
		"send_at",
	).Find(
		&entities,
	).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &entities
}

// 通知コード、トークンより取得
func (manager *NotificationMessages) Select(
	messageCode string,
) *database.NotificationMessages {

	var entity database.NotificationMessages

	err := Db.Where(
		"message_code = ?",
		messageCode,
	).First(
		&entity,
	).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &entity
}

// 更新＆作成.
func (manager *NotificationMessages) UpdateInsert(
	messageCode *string,
	notificationCode string,
	sendAt string,
	platform string,
	message string,
	sendCondition string,
) (bool, *string) {

	Db.Begin()

	nowAt := time.Now()

	// メッセージコードの指定がある場合
	if messageCode != nil {
		entity := manager.Select(*messageCode)
		if entity == nil {
			Db.Rollback()
			return false, nil
		}

		// 更新
		entity.SendAt = sendAt
		entity.Platform = platform
		entity.Message = message
		entity.SendCondition = sendCondition
		entity.UpdateAt = nowAt
		err := Db.Update(entity).Error
		if err != nil {
			log.Println(err)
			Db.Rollback()
			return false, nil
		}

	} else {

		// メッセージコード取得
		messageCode := uuid.NewV4().String()

		// 追加
		insertEntity := database.NotificationMessages{
			MessageCode:      messageCode,
			NotificationCode: notificationCode,
			SendAt:           sendAt,
			Platform:         platform,
			Message:          message,
			SendCondition:    sendCondition,
			IsSend:           constants.SendTypeNone,
			UpdateAt:         nowAt,
			CreateAt:         nowAt,
		}
		err := Db.Create(&insertEntity).Error
		if err != nil {
			log.Println(err)
			Db.Rollback()
			return false, nil
		}
	}

	Db.Commit()

	return true, messageCode
}

// 送信フラグを送信済みに更新.
func (manager *NotificationMessages) UpdateSendFlag(
	messageCode string,
) bool {

	Db.Begin()

	nowAt := time.Now()

	// メッセージコードの指定がある場合
	entity := manager.Select(messageCode)
	if entity == nil {
		Db.Rollback()
		return false
	}

	// 更新
	entity.IsSend = constants.SendTypeSent
	entity.UpdateAt = nowAt
	err := Db.Update(entity).Error
	if err != nil {
		log.Println(err)
		Db.Rollback()
		return false
	}

	Db.Commit()

	return true
}
