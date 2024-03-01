package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omericoramos/vagas-de-emprego/schema"
)

// @BasePath /api/v1

// @Summary Update opening
// @Description Update a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @param id query string true "Opening identification"
// @param opening body UpdateOpeningRequest true "Opening data to update"
// @success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {

	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	// verifica se os dados passados pelo contexto não estão vazios
	if err := request.validate(); err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")

	//  verifica se o di não esta vazioa, se estiver retorna o erro
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	//  cria uma estrutura da tabela do banco
	opening := schema.Opening{}

	// verifica se o objeto existe no banco de dados, se não existir retorna o erro
	if erro := db.First(&opening, id).Error; erro != nil {
		sendError(ctx, http.StatusNotFound,
			fmt.Sprintf("opening update with id: %s not found", id))
		return
	}

	// se os dados do request não for vazio subistitui pelo dado opening
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	// se der erro na atualização retorna o erro, caso contrario atualiza os dados
	if err := db.Save(&opening).Error; err != nil {

		logger.ErrorF("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucces(ctx, "update-opening", opening)
}
