package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseError(ctx *gin.Context, code int) {
	responseData := &ResponseData{
		Code:    code,
		Message: GetMessage(code),
		Data:    make(map[string]interface{}),
	}
	ctx.JSON(http.StatusOK, responseData)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	responseData := &ResponseData{
		Code:    ErrCodeSuccess,
		Message: GetMessage(ErrCodeSuccess),
		Data:    make(map[string]interface{}),
	}
	ctx.JSON(http.StatusOK, responseData)
}
