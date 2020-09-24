package controller

import (
	"hello/constant"
	"net/http"

	"github.com/gin-gonic/gin"
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
	setResponse(ctx, http.StatusOK, response)
}

func Error(ctx *gin.Context, code constant.ResponseCode, data map[string]interface{}) {

	response := response{
		Code:    code,
		Message: constant.GetCodeText(code),
		Data:    data,
	}
	ctx.Abort()
	setResponse(ctx, http.StatusOK, response)
}

func ErrorWithMessage(ctx *gin.Context, code constant.ResponseCode, message string, data map[string]interface{}) {

	response := response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	ctx.Abort()
	setResponse(ctx, http.StatusOK, response)
}

func setResponse(ctx *gin.Context, statusCode int, resp response) {
	ctx.Set("response", resp)
	ctx.JSON(statusCode, resp)
}

func NotFound(ctx *gin.Context) {
	response := response{
		Code:    constant.CODE_404,
		Message: "页面不存在",
		Data:    gin.H{},
	}
	ctx.Abort()
	setResponse(ctx, http.StatusNotFound, response)
}

func NoMethod(ctx *gin.Context) {
	response := response{
		Code:    constant.CODE_404,
		Message: "Method不存在",
		Data:    gin.H{},
	}
	ctx.Abort()
	setResponse(ctx, http.StatusMethodNotAllowed, response)
}

func StatusUnauthorized(ctx *gin.Context) {
	response := response{
		Code:    constant.USER_VERIFY_FAILD,
		Message: "认证失败",
		Data:    gin.H{},
	}
	ctx.Abort()
	setResponse(ctx, http.StatusUnauthorized, response)
}

func StatusInternalServerError(ctx *gin.Context) {
	ctx.Status(http.StatusInternalServerError)
	response := response{
		Code:    constant.CODE_500,
		Message: "服务器内部错误",
		Data:    gin.H{},
	}
	ctx.Abort()
	setResponse(ctx, http.StatusInternalServerError, response)
}
