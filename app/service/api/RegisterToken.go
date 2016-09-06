package api

import (
	"../../entity/request"
	"github.com/ant0ine/go-json-rest/rest"
)

// トークン登録サービス.
type RegisterToken struct {
}

// インスタンス.
var instanceRegisterToken *RegisterToken

// インスタンス取得.
func GetRegisterToken() *RegisterToken {
	if instanceRegisterToken == nil {
		instanceRegisterToken = &RegisterToken{}
	}
	return instanceRegisterToken
}

// 受信.
func (service *RegisterToken) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckAppToken(w, req)
	if !isSuccess {
		return
	}
	params := request.RegisterToken{}
	params.Convert(requestEntity.Params)

	//// 追加
	//isSuccess = service.locationsManager.DeleteInsert(
	//	requestEntity.Status.UserCode,
	//	requestEntity.Status.UserName,
	//	params,
	//)
	//if !isSuccess {
	//	w.WriteJson(util.ErrorResponseEntity(http.StatusInternalServerError, constants.MessageE4005))
	//	return
	//}

	//// レスポンス作成
	//metaEntity := response.LocationSend{}
	// レスポンス返却
	//w.WriteJson(util.OKResponseEntity(metaEntity))
}
