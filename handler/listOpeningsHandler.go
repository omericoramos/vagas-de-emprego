package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omericoramos/vagas-de-emprego/schema"
)

func ListOpeningsHandler(ctx *gin.Context) {

	openings := []schema.Opening{}

	if err := db.Find(&openings).Error; err != nil {

		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return

	}

	sendSucces(ctx, "list-openings", openings)
}
