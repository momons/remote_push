package util

import (
	"../entity/response"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

// レスポンスをオブジェクトとして取得.
func RequestEntity(entity interface{}, req *rest.Request) bool {
	// パラメータ取得
	err := req.DecodeJsonPayload(entity)
	if err != nil {
		entity = nil
		log.Println(err)
		return false
	}
	return true
}

// OKレスポンスを取得
func OKResponseEntity(metaEntity interface{}) response.Common {
	return response.Common{
		Status: response.Status{
			Code:    http.StatusOK,
			Message: "",
		},
		Meta: metaEntity,
	}
}

// エラーレスポンスを取得.
func ErrorResponseEntity(code int, message string) response.Common {
	return response.Common{
		Status: response.Status{
			Code:    code,
			Message: message,
		},
	}
}
