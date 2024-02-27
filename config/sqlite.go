package config

import (
	"os"

	"github.com/omericoramos/vagas-de-emprego/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {

	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// checa se o arquivo database existe
	_, err := os.Stat(dbPath)

	// se não existir tenta criar a pasta e o arquivo
	if os.IsNotExist(err) {

		logger.Info("Não existe o arquivo database, criando...")

		// cria a pasta db
		err = os.MkdirAll("./db", os.ModePerm)

		// se de erro printa e retorna o erro
		if err != nil {

			logger.ErrorF("Erro ao criar a pasta db: %v", err)
			return nil, err
		}

		// cria o arquivo database
		file, err := os.Create(dbPath)

		// se der erro printa e retorna o erro
		if err != nil {
			logger.ErrorF("Erro ao tentar criar o arquivo database: %v", err)
			return nil, err

		}

		// fecha o arquivo que foi criado
		file.Close()
	}

	// inicializando a conexão com o banco de dados
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	// se ouver erro printa e retorna o erro
	if err != nil {
		logger.ErrorF("Erro ao conectar no banco: %v", err)
		return nil, err
	}

	// roda as migrates
	err = db.AutoMigrate(&schema.Opening{})

	// se ouver erro printa e retorna o erro
	if err != nil {
		logger.ErrorF("Erro ao rodar as migrates: %v", err)
		return nil, err
	}

	return db, nil
}
