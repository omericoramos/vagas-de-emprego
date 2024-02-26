package router

import (
	"github.com/gin-gonic/gin"
	"github.com/omericoramos/vagas-de-emprego/handler"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/vagas", handler.ListOpeningsHandler)

		v1.GET("/vaga", handler.ShowOpeningHandler)

		v1.POST("/vaga", handler.CreateOpeningHandler)

		v1.DELETE("/vaga", handler.DeleteOpeningHandler)

		v1.PUT("/vaga", handler.UpdateOpeningHandler)
	}

}
