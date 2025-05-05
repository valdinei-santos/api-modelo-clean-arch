package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"
	create "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/usecases/create"
	get "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/usecases/get"
	getAll "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/usecases/get-all"
)

// Create - Controlador para criar um produto
func Create(log logger.Logger, ctx *gin.Context, useCase create.IUsecase) {
	log.Debug("Entrou controller.Get")
	var input *dto.Request
	err := json.NewDecoder(ctx.Request.Body).Decode(&input)
	if err != nil {
		log.Error(err.Error(), "mtd", "json.NewDecoder")
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, dataJErro)
		log.Info("### Finished ERROR", "status_code", http.StatusInternalServerError)
		return
	}

	resp, err := useCase.Execute(input)
	if err != nil {
		log.Error(err.Error(), "mtd", "useCase.Execute")
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, dataJErro)
		log.Info("### Finished ERROR", "status_code", http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, resp)
	log.Info("### Finished OK", "status_code", http.StatusOK)
	return
}

// Get - Controlador para obter um produto por ID
func Get(log logger.Logger, ctx *gin.Context, useCase get.IUsecase) {
	log.Debug("Entrou controller.Get")
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Error(err.Error(), "mtd", "strconv.Atoi")
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    "ID inválido",
		}
		ctx.JSON(http.StatusBadRequest, dataJErro)
		log.Info("### Finished ERROR", "status_code", http.StatusInternalServerError)
		return
	}

	log.Debug("ID: " + idParam)
	resp, err := useCase.Execute(id)
	if err != nil {
		log.Error(err.Error(), "mtd", "useCase.Execute")
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, dataJErro)
		log.Info("### Finished ERROR", "status_code", http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, resp)
	log.Info("### Finished OK", "status_code", http.StatusOK)
	return
}

// GetAll - Controlador para obter todos os produtos
func GetAll(log logger.Logger, ctx *gin.Context, useCase getAll.IUsecase) {
	log.Debug("Entrou controller.GetAll")
	resp, err := useCase.Execute()
	if err != nil {
		log.Error(err.Error(), "mtd", "useCase.Execute")
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, dataJErro)
		log.Info("### Finished ERROR", "status_code", http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, resp)
	log.Info("### Finished OK", "status_code", http.StatusOK)
	return
}
