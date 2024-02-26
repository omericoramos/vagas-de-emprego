package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Criando uma nova vaga",
	})
}

func ShowOpeningHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Detalhes da vaga cadastrada",
	})
}

func ListOpeningsHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Listando todas as vagas aqui",
	})
}

func UpdateOpeningHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Atualizando os dados da vaga",
	})
}

func DeleteOpeningHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Deletando uma vaga",
	})
}
