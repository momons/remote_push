package main

import (
	"./manager"
	"./service/api"
	"flag"
	"os"
)

// コマンド情報
// APIポート番号
var apiPort int

// ポーリングポート番号
var pollingPort int

// 終了コード
var exitCode = 0

// メイン
func main() {

	// セットアップ
	isSuccess := setup()
	if !isSuccess {
		exitCode = 1
	}

	os.Exit(exitCode)
}

// セットアップ
func setup() bool {

	// コマンドライン取得
	setupCommand()

	// データベース
	isSuccess := manager.SetupDatabase()
	if !isSuccess {
		return false
	}

	// API
	apiService := api.GetApiService()
	apiService.Start(apiPort)

	return true
}

// コマンドライン設定
func setupCommand() {
	// APIポート
	flag.IntVar(&apiPort, "apiport", 9005, "APIポートを指定して下さい。")

	flag.Parse()
}
