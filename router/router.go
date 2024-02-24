package router

// se eu não der um nome por padrão o go pegara o ultimo nome apos a ultima barra
import "github.com/gin-gonic/gin"

// para a função ser exportada deve ser iniciada com a letra maiuscula
func Initialize() {

	// Inicializando o router utiliozando as configurações padrão do gin
	router := gin.Default()

	// Definindo um routa GET
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Rodando a API na porta 8000
	router.Run(":8000") // listen and serve on 0.0.0.0:8080
}
