package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type response struct {
	Data interface{} `json:"data"`
}

type msgResponse struct{
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
// Success escribe una respuesta exitosa devolviendo la estructura
func Success(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, response{
		Data: data,
	})
}

// SuccessMsg escribe una respuesta exitosa sin estructura
func SuccessMsg(ctx *gin.Context, status int, msg string) {
	ctx.JSON(status, errorResponse{
		Message: msg,
		Status: status,
		Code: http.StatusText(status),
	})
}

// Failure escribe una respuesta fallida
func Failure(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, errorResponse{
		Message: err.Error(),
		Status:  status,
		Code:    http.StatusText(status),
	})
}
