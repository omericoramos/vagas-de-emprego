package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/vagas", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Lsitando todas as vagas",
			})
		})

		v1.GET("/vaga", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Exibindo os detalhes da vaga ",
			})
		})

		v1.POST("/vaga", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Criando uma nova vaga",
			})
		})

		v1.DELETE("/vaga", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Deletando uma vaga",
			})
		})
		v1.PUT("/vaga", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Atualizando uma vaga",
			})
		})
	}

}
