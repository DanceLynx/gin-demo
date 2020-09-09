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
	response := response{
		Code:    constant.CODE_404,
		Message: "页面不存在",
		Data:    gin.H{},
	}
	ctx.JSON(http.StatusNotFound, response)
}

func NoMethod(ctx *gin.Context) {
	response := response{
		Code:    constant.CODE_404,
		Message: "Method不存在",
		Data:    gin.H{},
	}
	ctx.JSON(http.StatusMethodNotAllowed, response)
}

func StatusUnauthorized(ctx *gin.Context) {
	response := response{
		Code:    constant.USER_VERIFY_FAILD,
		Message: "认证失败",
		Data:    gin.H{},
	}
	ctx.JSON(http.StatusUnauthorized, response)
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