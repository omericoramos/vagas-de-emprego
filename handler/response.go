package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omericoramos/vagas-de-emprego/schema"
)

func sendError(ctx *gin.Context, code int, message string) {

	ctx.Header("Content-type", "application/jason")
	ctx.JSON(code, gin.H{
		"message:":   message,
		"errorCode:": code,
	})

}

func sendSucces(ctx *gin.Context, op string, data interface{}) {

	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Operation from handler: %s successfull", op),
		"data":    data,
	})

}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string                 `json:"message"`
	Data    schema.OpeningResponse `json:"data"`
}
