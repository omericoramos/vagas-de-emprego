package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/omericoramos/vagas-de-emprego/docs"
	"github.com/omericoramos/vagas-de-emprego/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	// Inicializar o handler (manipulador)
	handler.InitializeHandlre()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/vagas", handler.ListOpeningsHandler)

		v1.GET("/vaga", handler.ShowOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.DELETE("/vaga", handler.DeleteOpeningHandler)

		v1.PATCH("/vaga", handler.UpdateOpeningHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
