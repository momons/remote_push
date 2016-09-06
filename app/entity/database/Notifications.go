package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 通知マスタテーブルEntity.
type Notifications struct {
	// ID.
	Id int64 `db:"id"`
	// 通知コード.
	NotificationCode string `db:"notification_code"`
	// 通知名.
	NotificationName string `db:"notification_name"`
	// pemファイル名(iOS).
	PemFileName string `db:"pem_file_name"`
	// APIキー(Android).
	apiKey string `db:"api_key"`
	// 更新日時.
	UpdateAt time.Time `db:"update_at"`
	// 作成日時.
	CreateAt time.Time `db:"create_at"`
}

// セットアップ.
func SetupNotifications(db *gorm.DB) {
	db.AutoMigrate(&Notifications{})
}
