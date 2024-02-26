package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Listando todas as vagas aqui",
	})
}
