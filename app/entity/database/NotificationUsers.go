package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 通知ユーザテーブルEntity.
type NotificationUsers struct {
	// ID.
	Id int64 `db:"id"`
	// 通知コード.
	NotificationCode string `db:"notification_code"`
	// 通知トークン.
	NotificationToken string `db:"notification_token"`
	// プラットフォーム.
	platform string `db:"platform"`
	// カスタムパラメータ.
	custom_params string `db:"custom_params"`
	// 更新日時.
	UpdateAt time.Time `db:"update_at"`
	// 作成日時.
	CreateAt time.Time `db:"create_at"`
}

// セットアップ.
func SetupNotificationUsers(db *gorm.DB) {
	db.AutoMigrate(&NotificationUsers{})
}
