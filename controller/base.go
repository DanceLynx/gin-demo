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

func Error(ctx *gin.Context, code constant.ResponseCode) {

	response := response{
		Code:    code,
		Message: constant.GetCodeText(code),
		Data:    gin.H{},
	}
	setResponse(ctx, http.StatusOK, response)
}

func ErrorWithMessage(ctx *gin.Context, code constant.ResponseCode, message string) {

	response := response{
		Code:    code,
		Message: message,
		Data:    gin.H{},
	}
	setResponse(ctx, http.StatusOK, response)
}

func setResponse(ctx *gin.Context, statusCode int, resp response) {
	ctx.Set("response", resp)
	ctx.JSON(statusCode, resp)
}

//NOTFOUND method not found action
func NOTFOUND(ctx *gin.Context) {
	response := response{
		Code:    constant.CODE_404,
		Message: constant.GetCodeText(constant.CODE_404),
		Data:    gin.H{},
	}
	ctx.Set("response", response)
	ctx.JSON(http.StatusNotFound, response)
}

//StatusInternalServerError server 500 error
func StatusInternalServerError(ctx *gin.Context) {
	response := response{
		Code:    constant.CODE_500,
		Message: constant.GetCodeText(constant.CODE_500),
		Data:    gin.H{},
	}
	ctx.Set("response", response)
	ctx.JSON(http.StatusInternalServerError, response)
}
