package controller

import (
	"github.com/gin-gonic/gin"
	"hello/constant"
	"net/http"
)

type response struct {
	Code    constant.ResponseCode `json:"code"`
	Message string                `json:"message"`
	Data    interface{}           `json:"data"`
}

func Success(ctx *gin.Context, message string, data map[string]interface{}) {

	response := response{
		Code:    constant.SUCCESS,
		Message: message,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, response)
}

func Error(ctx *gin.Context, code constant.ResponseCode, message string, data map[string]interface{}) {

	response := response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, response)
}

func NotFound(ctx *gin.Context) {
	ctx.Status(http.StatusNotFound)
	response := response{
		Code:    constant.CODE_404,
		Message: "页面不存在",
		Data:    gin.H{},
	}
	ctx.JSON(http.StatusNotFound, response)
}

func StatusInternalServerError(ctx *gin.Context) {
	ctx.Status(http.StatusInternalServerError)
	response := response{
		Code:    constant.CODE_500,
		Message: "服务器内部错误",
		Data:    gin.H{},
	}
	ctx.JSON(http.StatusInternalServerError, response)
}
