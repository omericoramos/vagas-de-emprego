package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omericoramos/vagas-de-emprego/schema"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @param id query string true "Opening identification"
// @success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]
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
