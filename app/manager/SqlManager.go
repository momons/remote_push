package manager

import (
	"../constants"
	"../entity/database"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

// データベース.
var Db *gorm.DB

// データベース設定.
func SetupDatabase() bool {

	// 接続文字列
	connectStr := "sslmode=disable host=localhost" + " dbname=" + constants.DatabaseName + " user=" + constants.UserId + " password=" + constants.Password
	db, err := gorm.Open("postgres", connectStr)
	if err != nil {
		log.Println(err)
		return false
	}
	db.DB()

	// デバッグモード.
	db.LogMode(true)

	// 退避
	Db = db

	// テーブル設定.
	database.SetupNotifications(db)
	database.SetupNotificationUsers(db)
	database.SetupNotificationMessages(db)

	return true
}
