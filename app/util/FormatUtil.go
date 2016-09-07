package util

import (
	"log"
	"time"
)

const (
	// 日付フォーマット.
	DateFormatString = "20060102150405"
)

// 文字列を日付型に変換.
func DataFromString(dateString string) *time.Time {
	date, err := time.Parse(DateFormatString, dateString)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &date
}

// 日付を文字列に変換
func StringFromDate(date time.Time) string {
	return date.Format(DateFormatString)
}
