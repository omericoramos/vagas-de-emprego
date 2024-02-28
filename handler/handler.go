package handler

import (
	"github.com/omericoramos/vagas-de-emprego/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandlre() {

	logger = config.GetLogger("handler")
	db = config.GetSQLite()

}
