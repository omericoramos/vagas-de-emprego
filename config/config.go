package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {

	var err error

	db, err = InitializeSQLite()

	if err != nil {

		return err
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

// Inicializando o log de erro
func GetLogger(prefix string) *Logger {

	logger = NewLogger(prefix)
	return logger
}
