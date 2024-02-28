package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omericoramos/vagas-de-emprego/schema"
)

func DeleteOpeningHandler(ctx *gin.Context) {

	// pega o valor do id passado no contexto
	id := ctx.Query("id")

	// verifica se não esta avzio
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schema.Opening{}

	// verifica se o objeto realmente exite no banco de dados
	if erro := db.First(&opening, id).Error; erro != nil {

		sendError(ctx, http.StatusNotFound,
			fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	// deletando os dados do banco
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id %s", id))
		return
	}

	sendSucces(ctx, "delete-opening", opening)
}
