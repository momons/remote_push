package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 通知メッセージテーブルEntity.
type NotificationMessages struct {
	// ID.
	Id int64 `db:"id"`
	// メッセージコード.
	MessageCode string `db:"message_code"`
	// 通知コード.
	NotificationCode string `db:"notification_code"`
	// 送信日時.
	SendAt time.Time `db:"send_at"`
	// プラットフォーム.
	platform string `db:"platform"`
	// 送信メッセージ.
	message string `db:"message"`
	// 送信条件 (JSON).
	send_condition string `db:"send_condition"`
	// 送信済み.
	is_send string `db:"is_send"`
	// 更新日時.
	UpdateAt time.Time `db:"update_at"`
	// 作成日時.
	CreateAt time.Time `db:"create_at"`
}

// セットアップ.
func SetupNotificationMessages(db *gorm.DB) {
	db.AutoMigrate(&NotificationMessages{})
}
