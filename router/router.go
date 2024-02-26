package router

// se eu não der um nome por padrão o go pegara o ultimo nome apos a ultima barra
import "github.com/gin-gonic/gin"

// para a função ser exportada deve ser iniciada com a letra maiuscula
func Initialize() {

	// Inicializando o router utilizando as configurações padrão do gin
	router := gin.Default()

	// Inicializando as rotas
	initializeRoutes(router)

	// por padrão roda o server na porta 0.0.0.0:8080
	router.Run(":8000") // difinindo a porta 8000 para o server

}
