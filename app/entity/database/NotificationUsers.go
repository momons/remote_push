package database

import (
	"../../constants"
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
	Platform string `db:"platform"`
	// カスタムパラメータ.
	CustomParams string `db:"custom_params"`
	// 更新日時.
	UpdateAt time.Time `db:"update_at"`
	// 作成日時.
	CreateAt time.Time `db:"create_at"`
}

// セットアップ.
func SetupNotificationUsers(db *gorm.DB) {
	db.AutoMigrate(&NotificationUsers{})
}

// プラットフォームがiOSか
func (entity *NotificationUsers) IsiOS() bool {
	return entity.Platform == constants.PlatformTypeiOS
}

// プラットフォームがAndroidか
func (entity *NotificationUsers) IsAndroid() bool {
	return entity.Platform == constants.PlatformTypeAndroid
}
