package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omericoramos/vagas-de-emprego/schema"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @param request body CreateOpeningRequest true "Request body"
// @success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {

	request := CreateOpeningRequest{}

	// popula o resquest com os dados vindo do ctx utilizando ctx.BindJson
	ctx.BindJSON(&request)

	if err := request.validate(); err != nil {

		logger.ErrorF("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schema.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {

		logger.ErrorF("Erro ao criar a nova vaga %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
	}
	sendSucces(ctx, "create-opening", opening)
}
