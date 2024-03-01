package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omericoramos/vagas-de-emprego/schema"
)

// @BasePath /api/v1

// @Summary List opening
// @Description List a new job openings
// @Tags Openings
// @Accept json
// @Produce json
// @success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {

	openings := []schema.Opening{}

	if err := db.Find(&openings).Error; err != nil {

		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return

	}

	sendSucces(ctx, "list-openings", openings)
}
