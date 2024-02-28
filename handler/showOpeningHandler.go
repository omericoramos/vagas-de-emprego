package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omericoramos/vagas-de-emprego/schema"
)

func ShowOpeningHandler(ctx *gin.Context) {

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schema.Opening{}

	if err := db.First(&opening, id).Error; err != nil {

		sendError(ctx, http.StatusNotFound,
			fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	sendSucces(ctx, "show-opening", opening)
}
