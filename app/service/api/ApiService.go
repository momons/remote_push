package api

import (
	"../../constants"
	"../../entity/request"
	"../../util"
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

// APIサービス.
type ApiService struct {
	// API
	api *rest.Api
}

// インスタンス.
var instanceApiService *ApiService

// インスタンス取得.
func GetApiService() *ApiService {
	if instanceApiService == nil {
		instanceApiService = &ApiService{}
		instanceApiService.setup()
	}
	return instanceApiService
}

// APIサービスを設定.
func (service *ApiService) setup() bool {

	service.api = rest.NewApi()
	service.api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		// トークン取得
		&rest.Route{"POST", "/register_token", GetRegisterToken().Recive},
	)
	if err != nil {
		log.Fatal(err)
		return false
	}
	service.api.SetApp(router)

	return true
}

// スタート.
func (service *ApiService) Start(port int) {
	// ポート番号を文字列化
	portStr := fmt.Sprintf("%d", port)

	log.Printf("API server started. port=%s", portStr)
	log.Fatal(http.ListenAndServe(":"+portStr, service.api.MakeHandler()))
}

// リクエスト取得＆アプリトークンチェック
func GetRequestAndCheckAppToken(
	w rest.ResponseWriter,
	req *rest.Request,
) (bool, *request.Common) {

	// リクエスト取得
	requestEntity := request.Common{}
	isSuccess := util.RequestEntity(&requestEntity, req)
	if !isSuccess {
		w.WriteJson(util.ErrorResponseEntity(http.StatusBadRequest, constants.MessageE0001))
		return false, nil
	}

	// アプリトークンチェック
	if !requestEntity.Status.IsValidAppToken() {
		w.WriteJson(util.ErrorResponseEntity(http.StatusBadRequest, constants.MessageE0003))
		return false, &requestEntity
	}

	return true, &requestEntity
}
