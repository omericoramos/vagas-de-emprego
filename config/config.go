package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {

	return nil
}

// Inicializando o log de erro
func GetLogger(prefix string) *Logger {

	logger = NewLogger(prefix)
	return logger
}
