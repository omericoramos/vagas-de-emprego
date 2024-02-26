package main

import (
	"github.com/omericoramos/vagas-de-emprego/config"
	"github.com/omericoramos/vagas-de-emprego/router"
)

var (
	logger *config.Logger
)

func main() {

	// Inicializando as configurações
	err := config.Init()

	logger = config.GetLogger("main")

	if err != nil {
		logger.ErrorF("Erro na inicialização da config: %s", err)
		return
	}

	// para a função ser exportada deve ser iniciada com a letra maiuscula
	router.Initialize()
}
